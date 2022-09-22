# netlink

# 概述
> 参考：[GitHub 项目](https://github.com/vishvananda/netlink)

netlink 包为 go 提供了一个简单的 netlink 库。Netlink是linux中的用户空间程序用来与内核进行通信的界面。它可以用于添加和删除接口，设置ip地址和路由以及配置ipsec。Netlink通信需要提升的权限，因此在大多数情况下，此代码需要以root用户身份运行。由于低级netlink消息充其量是难以理解的，因此该库试图提供一个以iproute2提供的CLI为松散建模的api。ip链接添加之类的操作将通过类似名称的函数 (例如AddLink()) 来完成。该库的生命开始于docker/libcontainer中的netlink功能分支，但经过大量重写以提高可测试性，性能并添加ipsec xfrm处理等新功能。