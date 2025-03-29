### **📌 2周 Go 业务开发学习计划（适合工作日晚上 + 周末学习）**  
目标：**2周内上手 Go 后端业务开发，具备 API 开发、数据库操作、缓存、部署能力**  

---

## **📆 第一周：核心技能学习**
**工作日（晚上2-3小时）**：学习核心技术  
**周末（整块时间）**：项目实战  

### **📌 第1天：Go Web 框架 Gin**
✅ **目标**：掌握 Gin 进行 API 开发  
- 了解 Gin 基本概念（路由、中间件、请求处理）  
- 编写一个简单的 `Hello World` API  
- 实现 RESTful API（增删改查）   

**练习任务**：
- 创建 `/users` 相关 API（`GET /users`，`POST /users`，`DELETE /users/:id`）  
- 学习 JSON 解析（`c.BindJSON()`）  

📖 **推荐学习资料**：
- [Gin 官方文档](https://gin-gonic.com/docs/)  
- [示例代码](https://github.com/gin-gonic/examples)  

---

### **📌 第2天：数据库操作 GORM**
✅ **目标**：学会使用 GORM 进行数据库操作  
- 连接 MySQL / PostgreSQL / SQLite  
- 定义数据模型（结构体 -> 数据库表）  
- 增删改查（CRUD）  

**练习任务**：
- 设计 `users` 数据表（`ID`，`Name`，`Email`）  
- 编写 `/users` API，并存储到数据库  

📖 **推荐学习资料**：
- [GORM 官方文档](https://gorm.io/docs/)  
- [GORM 使用示例](https://github.com/go-gorm/gorm)  

---

### **📌 第3天：用户认证（JWT）**
✅ **目标**：实现用户登录 & Token 认证  
- JWT 介绍（JSON Web Token）  
- 使用 `github.com/golang-jwt/jwt` 生成 & 解析 Token  
- 在 Gin 中实现 **登录 & 鉴权** 中间件  

**练习任务**：
- 编写 `POST /login` 登录 API（返回 Token）  
- 保护 `/users` API，要求携带 Token 才能访问  

📖 **推荐学习资料**：
- [JWT 官方文档](https://jwt.io/)  

---

### **📌 第4天：Redis 缓存**
✅ **目标**：使用 Redis 进行数据缓存  
- 安装 Redis 并连接 Go 应用  
- 存取数据（`Set/Get` 操作）  
- 结合 Redis 缓存 API 数据，提升查询性能  

**练习任务**：
- 给 `/users` API 添加缓存（先查缓存，再查数据库）  

📖 **推荐学习资料**：
- [Go-Redis 官方文档](https://github.com/go-redis/redis)  

---

### **📌 第5天：Gin + GORM + Redis 业务整合**
✅ **目标**：将前面学的内容整合，构建一个简单的 API 服务  
- 结合 Gin + GORM + Redis，实现完整的用户管理功能  
- 增强日志（使用 `logrus`）  
- 处理 API 异常  

📖 **推荐学习资料**：
- [logrus 官方文档](https://github.com/sirupsen/logrus)  

---

### **📌 第6-7天（周末）：项目实战**
✅ **目标**：实现一个 **短链接服务**（TinyURL 类似）  
**需求**：
- 用户输入长链接，生成短链接  
- 短链接重定向到原始链接  
- 统计访问次数  

📖 **推荐示例**：
- [短链接服务实现](https://github.com/eddycjy/go-short-url)  

---

## **📆 第二周：深入业务 & 实战项目**
**工作日（晚上2-3小时）**：学习新技术 & 代码优化  
**周末（整块时间）**：深度开发业务项目  

### **📌 第8天：任务队列（异步处理）**
✅ **目标**：学会用 RabbitMQ / Kafka 处理异步任务  
- 了解消息队列的作用  
- 使用 `go-redis` 的 `List` 作为简单任务队列  
- 生产者 / 消费者模式  

📖 **推荐学习资料**：
- [RabbitMQ Go 客户端](https://github.com/rabbitmq/amqp091-go)  

---

### **📌 第9天：定时任务**
✅ **目标**：学习如何在 Go 中实现定时任务  
- `cron` 表达式 & Go 任务调度库  
- 实现一个每日任务（如：清理过期数据）  

📖 **推荐学习资料**：
- [Go cron 库](https://github.com/robfig/cron)  

---

### **📌 第10天：API 性能优化**
✅ **目标**：提升 API 速度 & 限流  
- 使用 `gin.Limiter` 进行限流（防止 DDoS）  
- Redis 作为 **请求计数器**（实现访问频率限制）  

📖 **推荐学习资料**：
- [GoRate 限流](https://github.com/uber-go/ratelimit)  

---

### **📌 第11天：Docker & 部署**
✅ **目标**：学会使用 Docker 部署 Go 服务  
- 编写 `Dockerfile` 打包 Go 应用  
- 使用 `docker-compose` 启动 Go + Redis + MySQL  

📖 **推荐学习资料**：
- [Docker + Go 教程](https://docs.docker.com/samples/golang/)  

---

### **📌 第12-13天（周末）：业务系统实战**
✅ **目标**：独立开发一个完整的 **订单管理系统**  
- 用户下单、订单支付、订单查询  
- 结合 JWT、Redis、GORM、消息队列  
- **重点**：设计数据库表 & API 逻辑  

📖 **推荐学习资料**：
- [完整订单系统示例](https://github.com/go-ecommerce/example)  

---

### **📌 第14天（总结 & 复盘）**
✅ **目标**：总结学习内容，优化代码  
- 代码重构：优化 API 结构、日志、异常处理  
- 业务总结：整理 API 文档，思考可以优化的点  

---

## **📌 2周学习节奏总结**
| **时间**     | **学习内容**                 | **练习任务**                     |
|------------|--------------------|----------------------|
| **第1周**  | Gin、GORM、JWT、Redis | 开发短链接服务 |
| **第2周**  | 队列、限流、Docker、部署 | 开发订单管理系统 |

---

## **🎯 结论**
**2周后，你将具备：**  
✅ Gin + GORM 开发 RESTful API 能力  
✅ 使用 Redis 进行缓存优化  
✅ JWT 实现用户认证 & 限流优化  
✅ RabbitMQ / Redis 队列处理异步任务  
✅ Docker 部署 Go 服务  
✅ 能够开发完整的业务系统 🚀  

这样，你不仅能 **上手开发业务**，还能写出 **高质量、可扩展的 Go 后端代码！** 💪🔥