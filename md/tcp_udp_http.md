# tcp udp http

## 1. tcp

### 1. 三次握手



#### 1. 为什么是三次握手，而不是是二次？

*主要有两点*

- 双方都准备好了
- 协商得知对方的序列号（因此三次握手后的序列号就是C/S交互要依赖的序列号 seq）





#### 2. 为什么seq序列号不是从0开始的？

**seq** 的取值范围是 0 ～ 2^32

主要是为了减少重复的概率，防止包重复发送。





https://www.bilibili.com/video/BV13z4y1D7Pc?from=search&seid=8212530173527907065

https://mp.weixin.qq.com/s/GQTVytNus8EcoGEkOSTUog



### 2. 四次挥手



### 3. 2msl的time_wait

2MSL Maximum Segment Life
主要是为了保证第三次FIN请求后，在第四次ACK确认后，防止对方没有收到。
2MSL就是2次发送FIN-ACK的时间周期，如果对方没有重发第三次FIN，就表示对方已经收到了ACK
就可以放心进入关闭状态啦。


## 2. udp



UDP的应用场景：直播，QQ，DNS



## 3. http

### 1. http状态码

- 200  // 请求成功 GET｜POST
- 201  // 请求成功，并创建了资源

---
- 301 // 永久重定向
- 302 // 临时重定向

---
- 400 // Bad Request 请求参数错误
- 401 // 未认证  
- 403 // 权限拒绝
- 404 // 资源未找到
- 405 // 请求方式不允许

---
- 500 // Internal Server Error 
- 502 // bad gateway 网关错误，后端服务异常
- 504 // gateway time-out 网关超时，后端服务挂啦

https://baike.baidu.com/item/HTTP状态码/5053660?fr=aladdin#5_4

### 2. http 2.0

- 多路复用
- 服务端推送
- header压缩

https://www.zhihu.com/question/34074946





