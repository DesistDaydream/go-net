package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
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

// UDP 伪头部（用于计算 UDP 校验和）
func udpChecksum(srcIP, dstIP net.IP, udpPacket []byte) uint16 {
	pseudoHeader := make([]byte, 12)
	copy(pseudoHeader[0:4], srcIP.To4())
	copy(pseudoHeader[4:8], dstIP.To4())
	pseudoHeader[8] = 0
	pseudoHeader[9] = syscall.IPPROTO_UDP
	binary.BigEndian.PutUint16(pseudoHeader[10:12], uint16(len(udpPacket)))

	data := append(pseudoHeader, udpPacket...)
	return checksum(data)
}

// 发送源端口为 0 的 UDP 包。验证 0 端口号是否可以正常在网络上流转
func main() {
	// 创建原始套接字（需要 root 权限）
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		fmt.Println("创建原始套接字失败（需要 root 权限）:", err)
		return
	}
	defer syscall.Close(fd)

	// 设置 IP_HDRINCL 选项，告诉内核我们自己构造 IP 头
	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		fmt.Println("设置 IP_HDRINCL 失败:", err)
		return
	}

	// 源和目标配置
	srcIP := net.ParseIP("10.10.4.90").To4()  // 修改为你的本机 IP
	dstIP := net.ParseIP("10.10.11.54").To4() // 修改为目标 IP
	dstPort := uint16(8888)
	payload := []byte("Hello")

	// 构造 UDP 头部（8字节）
	udpHeader := make([]byte, 8)
	binary.BigEndian.PutUint16(udpHeader[0:2], 0)                      // 源端口 = 0
	binary.BigEndian.PutUint16(udpHeader[2:4], dstPort)                // 目标端口
	binary.BigEndian.PutUint16(udpHeader[4:6], uint16(8+len(payload))) // UDP 长度
	binary.BigEndian.PutUint16(udpHeader[6:8], 0)                      // 校验和占位

	// 组合 UDP 数据包
	udpPacket := append(udpHeader, payload...)

	// 计算 UDP 校验和
	csum := udpChecksum(srcIP, dstIP, udpPacket)
	binary.BigEndian.PutUint16(udpPacket[6:8], csum)

	// 构造 IP 头部（20字节，无选项）
	ipHeader := make([]byte, 20)
	ipHeader[0] = 0x45                                                   // 版本(4) + 头长度(5)
	ipHeader[1] = 0                                                      // TOS
	binary.BigEndian.PutUint16(ipHeader[2:4], uint16(20+len(udpPacket))) // 总长度
	binary.BigEndian.PutUint16(ipHeader[4:6], 54321)                     // ID
	binary.BigEndian.PutUint16(ipHeader[6:8], 0)                         // Flags + Fragment offset
	ipHeader[8] = 64                                                     // TTL
	ipHeader[9] = syscall.IPPROTO_UDP                                    // 协议
	binary.BigEndian.PutUint16(ipHeader[10:12], 0)                       // 校验和占位
	copy(ipHeader[12:16], srcIP)                                         // 源 IP
	copy(ipHeader[16:20], dstIP)                                         // 目标 IP

	// 计算 IP 校验和
	binary.BigEndian.PutUint16(ipHeader[10:12], checksum(ipHeader))

	// 组合完整数据包
	packet := append(ipHeader, udpPacket...)

	// 发送
	addr := syscall.SockaddrInet4{
		Port: 0,
	}
	copy(addr.Addr[:], dstIP)

	err = syscall.Sendto(fd, packet, 0, &addr)
	if err != nil {
		fmt.Println("发送失败:", err)
		return
	}

	fmt.Printf("已发送源端口为 0 的 UDP 包到 %s:%d\n", dstIP, dstPort)
}
