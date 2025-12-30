package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"syscall"

	"github.com/sirupsen/logrus"
)

// 计算校验和
func checksum(data []byte) uint16 {
	sum := uint32(0)
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(binary.BigEndian.Uint16(data[i:]))
	}
	if len(data)%2 == 1 {
		sum += uint32(data[len(data)-1]) << 8
	}
	for sum > 0xffff {
		sum = (sum & 0xffff) + (sum >> 16)
	}
	return ^uint16(sum)
}

// 伪头部校验和（UDP/TCP 通用）
func transportChecksum(srcIP, dstIP []byte, protocol uint8, transportPacket []byte) uint16 {
	pseudoHeader := make([]byte, 12)
	copy(pseudoHeader[0:4], srcIP)
	copy(pseudoHeader[4:8], dstIP)
	pseudoHeader[8] = 0
	pseudoHeader[9] = protocol
	binary.BigEndian.PutUint16(pseudoHeader[10:12], uint16(len(transportPacket)))

	data := append(pseudoHeader, transportPacket...)
	return checksum(data)
}

// 构造 IP 头部
func buildIPHeader(srcIP, dstIP []byte, protocol uint8, payloadLen int) []byte {
	ipHeader := make([]byte, 20)
	ipHeader[0] = 0x45                                               // 版本(4) + 头长度(5)
	ipHeader[1] = 0                                                  // TOS
	binary.BigEndian.PutUint16(ipHeader[2:4], uint16(20+payloadLen)) // 总长度
	binary.BigEndian.PutUint16(ipHeader[4:6], 54321)                 // ID
	binary.BigEndian.PutUint16(ipHeader[6:8], 0)                     // Flags + Fragment offset
	ipHeader[8] = 64                                                 // TTL
	ipHeader[9] = protocol                                           // 协议
	binary.BigEndian.PutUint16(ipHeader[10:12], 0)                   // 校验和占位
	copy(ipHeader[12:16], srcIP)                                     // 源 IP
	copy(ipHeader[16:20], dstIP)                                     // 目标 IP

	// 计算 IP 校验和
	binary.BigEndian.PutUint16(ipHeader[10:12], checksum(ipHeader))
	return ipHeader
}

// 构造 UDP 数据包
func buildUDPPacket(srcIP, dstIP []byte, srcPort, dstPort uint16, payload []byte) []byte {
	// 构造 UDP 头部（8字节）
	udpHeader := make([]byte, 8)
	binary.BigEndian.PutUint16(udpHeader[0:2], srcPort)                // 源端口
	binary.BigEndian.PutUint16(udpHeader[2:4], dstPort)                // 目标端口
	binary.BigEndian.PutUint16(udpHeader[4:6], uint16(8+len(payload))) // UDP 长度
	binary.BigEndian.PutUint16(udpHeader[6:8], 0)                      // 校验和占位

	// 组合 UDP 数据包
	udpPacket := append(udpHeader, payload...)

	// 计算 UDP 校验和
	csum := transportChecksum(srcIP, dstIP, syscall.IPPROTO_UDP, udpPacket)
	binary.BigEndian.PutUint16(udpPacket[6:8], csum)

	// 构造 IP 头部
	ipHeader := buildIPHeader(srcIP, dstIP, syscall.IPPROTO_UDP, len(udpPacket))

	// 组合完整数据包
	return append(ipHeader, udpPacket...)
}

// 构造 TCP 数据包
func buildTCPPacket(srcIP, dstIP []byte, srcPort, dstPort uint16, payload []byte) []byte {
	// 构造 TCP 头部（20字节，无选项）
	tcpHeader := make([]byte, 20)
	binary.BigEndian.PutUint16(tcpHeader[0:2], srcPort) // 源端口
	binary.BigEndian.PutUint16(tcpHeader[2:4], dstPort) // 目标端口
	binary.BigEndian.PutUint32(tcpHeader[4:8], 0)       // 序列号
	binary.BigEndian.PutUint32(tcpHeader[8:12], 0)      // 确认号
	tcpHeader[12] = 0x50                                // 数据偏移(5) + 保留位
	tcpHeader[13] = 0x02                                // 标志位 (SYN)
	// tcpHeader[13] = 0xff                               // 标志位 (SYN)
	binary.BigEndian.PutUint16(tcpHeader[14:16], 8192) // 窗口大小
	binary.BigEndian.PutUint16(tcpHeader[16:18], 0)    // 校验和占位
	binary.BigEndian.PutUint16(tcpHeader[18:20], 0)    // 紧急指针

	// 组合 TCP 数据包
	tcpPacket := append(tcpHeader, payload...)

	// 计算 TCP 校验和
	csum := transportChecksum(srcIP, dstIP, syscall.IPPROTO_TCP, tcpPacket)
	binary.BigEndian.PutUint16(tcpPacket[16:18], csum)

	// 构造 IP 头部
	ipHeader := buildIPHeader(srcIP, dstIP, syscall.IPPROTO_TCP, len(tcpPacket))

	// 组合完整数据包
	return append(ipHeader, tcpPacket...)
}

// 发送原始数据包
func sendRawPacket(packet []byte, dstIP []byte) error {
	// 创建原始套接字（需要 root 权限）
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		return fmt.Errorf("创建原始套接字失败（需要 root 权限）: %v", err)
	}
	defer syscall.Close(fd)

	// 设置 IP_HDRINCL 选项
	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		return fmt.Errorf("设置 IP_HDRINCL 失败: %v", err)
	}

	// 设置目标地址
	addr := syscall.SockaddrInet4{
		Port: 0,
	}
	copy(addr.Addr[:], dstIP)

	err = syscall.Sendto(fd, packet, 0, &addr)
	if err != nil {
		return fmt.Errorf("发送失败: %v", err)
	}

	return nil
}

func main() {
	// 命令行参数
	protocol := flag.String("p", "udp", "协议类型: udp 或 tcp")
	srcIP := flag.String("src", "10.10.4.90", "源 IP 地址")
	dstIP := flag.String("dst", "10.10.11.54", "目标 IP 地址")
	srcPort := flag.Int("sport", 0, "源端口")
	dstPort := flag.Int("dport", 0, "目标端口")
	payload := flag.String("data", "DesistDaydream", "要发送的数据")

	flag.Parse()

	// 解析 IP 地址
	srcIPv4 := net.ParseIP(*srcIP).To4()
	dstIPv4 := net.ParseIP(*dstIP).To4()

	// 构造数据包
	var packet []byte
	switch *protocol {
	case "udp":
		packet = buildUDPPacket(srcIPv4, dstIPv4, uint16(*srcPort), uint16(*dstPort), []byte(*payload))
		fmt.Printf("构造 UDP 数据包: %s:%d -> %s:%d\n", srcIPv4, *srcPort, dstIPv4, *dstPort)
	case "tcp":
		packet = buildTCPPacket(srcIPv4, dstIPv4, uint16(*srcPort), uint16(*dstPort), []byte(*payload))
		fmt.Printf("构造 TCP 数据包 (SYN): %s:%d -> %s:%d\n", srcIPv4, *srcPort, dstIPv4, *dstPort)
	default:
		logrus.Fatalln("不支持的协议类型")
	}

	// 发送数据包
	err := sendRawPacket(packet, dstIPv4)
	if err != nil {
		fmt.Println("发送失败:", err)
		return
	}

	fmt.Printf("✓ 已成功发送 %d 字节\n", len(packet))
	fmt.Printf("Payload: %s\n", *payload)
}
