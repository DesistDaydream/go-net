package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-ping/ping"
	"github.com/spf13/pflag"
)

type pingFlags struct {
	Count   int
	Domains []string
	Output  string
	Network string
}

func (po *pingFlags) NewPingFlags() {
	pflag.IntVarP(&po.Count, "count", "c", 3, "ping 测试次数")
	pflag.StringSliceVarP(&po.Domains, "domains", "d", nil, "待测试域名列表，以逗号分隔，也可以多次使用 -d 标志")
	pflag.StringVarP(&po.Output, "output", "o", "text", "输出格式，可用的值有:{json|text}")
	pflag.StringVarP(&po.Network, "network", "n", "ip", "解析时使用的 IP 版本，默认自动选择，可用的值有:{ip|ip4|ip6}")
	pflag.Parse()
}

type pingResult struct {
	Domain     string        `json:"domain"`
	IP         net.IP        `json:"ip"`
	PacketLoss float64       `json:"packet_loss"`
	MinRTT     time.Duration `json:"min_rtt"`
	MaxRTT     time.Duration `json:"max_rtt"`
	AvgRTT     time.Duration `json:"avg_rtt"`
}

func NewPingResult(p *ping.Pinger) *pingResult {
	return &pingResult{
		Domain:     p.Statistics().Addr,
		IP:         p.Statistics().IPAddr.IP,
		PacketLoss: p.Statistics().PacketLoss,
		MinRTT:     p.Statistics().MinRtt,
		MaxRTT:     p.Statistics().MaxRtt,
		AvgRTT:     p.Statistics().AvgRtt,
	}
}

type Loggers []ping.Logger

func main() {
	pf := &pingFlags{}
	pf.NewPingFlags()

	for _, domain := range pf.Domains {
		// 实例化 Pinger
		pinger, err := ping.NewPinger(domain)
		if err != nil {
			panic(err)
		}
		// 设置一些执行 ping 时的参数
		pinger.SetPrivileged(false)
		pinger.SetNetwork(pf.Network)
		pinger.Count = pf.Count

		var wg sync.WaitGroup
		defer wg.Wait()

		wg.Add(1)

		go func() {
			defer wg.Done()
			// 执行 ping，并为 pinger 赋值
			err = pinger.Run()
			if err != nil {
				panic(err)
			}

			// 实例化 ping 的结果
			pingResult := NewPingResult(pinger)
			prJSON, err := json.Marshal(pingResult)
			if err != nil {
				panic(err)
			}

			// 输出结果
			switch pf.Output {
			case "json":
				fmt.Println(string(prJSON))
			case "text":
				fmt.Printf("拨测域名:%v\n"+
					"  解析地址:%v\n"+
					"  丢包百分比:%v\n"+
					"  最小RTT:%v\n"+
					"  最大RTT:%v\n"+
					"  平均RTT:%v\n",
					pingResult.Domain, pingResult.IP, pingResult.PacketLoss, pingResult.MinRTT, pingResult.MaxRTT, pingResult.AvgRTT)
			default:
				panic("输出格式错误")
			}
		}()
	}

}
