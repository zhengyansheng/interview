# channel

## 1.channel 类型

- 无缓冲区的
```go
intCh := make(chan int)
```
无缓冲区的 channel，如果只有一个主 goroutine, 只要往 channel 发送数据就会引起死锁 deadlock；
必须在往 channel 发送数据之前，在创建一个 goroutine 先从 channel 读取才行；

- 有缓冲区的
```go
intCh := make(chan int, 1)
```
有缓冲区的 channel，允许你发送指定容量的 channel 个数，如果超过就会死锁；

## 2.channel 底层如何实现阻塞的
互斥锁，因此能保证线程安全



## 3. channel发送和接收的本质是

>  All transfer of value on the go channels happens with the copy of value.
>
> 所有在go channels传输的值都是值的拷贝



## 4. 如果往已关闭的 channel 发送数据会怎样？为什么

- 死锁


## 参考
https://www.cnblogs.com/show58/p/12699083.html

