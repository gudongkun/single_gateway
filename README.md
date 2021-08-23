[## single u center service

###已经完成 ：
1. 消息队列broker引入
2. gomicro自动日志。
3. 分布式日志
4. 项目私有化
5. xorm日志入jeager
6. prometheus监控
7. gin网关
8. jwt
9.  角色管理 Casbin
       
###待完成：
1. jwt 和 cashbin编码完成，没有实际测
2. casbin可以写到 ucenter项目


### 特别注意

jaeger 不支持，gin.context 要用 c:= ctx.Request.Context()获取原生context