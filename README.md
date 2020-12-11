# Go 语言 Web 编程学习
参考：
* [GoWeb编程](https://github.com/astaxie/build-web-application-with-golang)
* [看云 GoWeb 编程](https://www.kancloud.cn/kancloud/web-application-with-golang)

go 使用 [net/http](https://pkg.go.dev/net/http) 标准库来实现基本的 web 功能

[hello_world](./hello_world/main.go) # 用于快速体验如何使用 go 来快速搭建一个 web 服务器  
[form](./form/main.go) # 描述网页表单的处理  

## 一般的上网过程  
浏览器本身是一个客户端，当你输入URL的时候，首先浏览器会去请求DNS服务器，通过DNS获取相应的域名对应的IP，然后通过IP地址找到IP对应的服务器后，要求建立TCP连接，等浏览器发送完HTTP Request（请求）包后，服务器接收到请求包之后才开始处理请求包，服务器调用自身服务，返回HTTP Response（响应）包；客户端收到来自服务器的响应后开始渲染这个Response包里的主体（body），等收到全部的内容随后断开与该服务器之间的TCP连接  