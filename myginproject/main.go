package main

import (
    "myginproject/api/ping"
    "myginproject/api/user"
    "myginproject/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    // 注册中间件
    r.Use(middleware.SimpleRecovery())
    
    // 注册各个模块的路由
    ping.RegisterRoutes(r)
    user.RegisterRoutes(r)
    // 后续添加其他模块的路由注册
    // product.RegisterRoutes(r)
    // order.RegisterRoutes(r)
    
    r.Run(":9000")
}
