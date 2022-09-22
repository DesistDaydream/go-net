package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {
	// 实例化一个 LinkAttrs,LinkAttrs 包含一个网络设备的绝大部分属性
	linkAttrs := netlink.NewLinkAttrs()
	// 设定 link 的名称
	linkAttrs.Name = "br0"
	// 将实例化的 LinkAttrs 信息赋值给 Bridge 结构体
	mybridge := &netlink.Bridge{LinkAttrs: linkAttrs}
	// 这里就算真正完成了一个网络设备的定义，netlink 库中包含多种网络设备结构体
	// 每种网络设备结构体都实现了 Link 接口
	// Link 接口只有两个方法，Attrs() 用来返回 LinkAttrs 结构体，Type() 用来返回该网络设备的类型。
	// 而对各种类型的网络设备实现增删改查的函数，其接受的参数就是 Link 接口类型
	// 所以 Link 接口的主要作用，就是用来区分不同类型的网络设备，以便可以在调用时统一。对网络设备的任何操作，都可以将 Link 接口作为参数互相传递。

	// 使用 Bridge 结构体的信息创建一个网络设备
	err := netlink.LinkAdd(mybridge)
	if err != nil {
		fmt.Printf("could not add %s: %v\n", linkAttrs.Name, err)
	}
	// eth1, _ := netlink.LinkByName("eth1")
	// netlink.LinkSetMaster(eth1, mybridge)
}
