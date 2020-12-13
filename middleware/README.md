# Middleware(中间件)
在 Web 编程中，Middleware(中间件) 一般是指在接收到用户的消息前，先进行一系列的预处理，然后再交给实际的 `http.Handler()` 去处理，处理结果可能还需要一系列的后续处理。

典型的中间件功能：
* 认证、鉴权
* 统计
* 等等...

![](https://raw.githubusercontent.com/DesistDaydream/PictureHosting/main/GoWeb/middleware.drawio)

加入我现在用 gin 框架注册两个路由
```go
	r.POST("/login", handler.LoginPost)
	r.GET("/order", handler.OrderGet)
```
如果想让客户端在登录之后才可以访问 /order 页面，那么就需要添加一个处理逻辑，比如：
```go
    r.POST("/login", handler.LoginPost)
    // 注册中间件，每次访问中间件之后的路由，都会执行一次中间件的逻辑
    r.Use(middleware.MiddleWare())
    r.GET("/order", handler.OrderGet)
```
这样，每次访问 /order 的时候，都会执行 `middleware.MiddleWare()` 中的逻辑，这个逻辑里可以进行判断，看看本次请求是否带有 TOKEN 或者 Session 之类的认证信息，并与本身数据库中的数据进行匹配，匹配成功才能继续访问 /order。