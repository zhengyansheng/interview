# CSP

> Communicating Sequential Processes
>
> 通信顺序进程





Do not communicate by sharing memory; instead, share memory by communicating.

不要通过共享内存来通信，而要通过通信来实现内存共享。



这就是 Go 的并发哲学，它依赖 CSP 模型，基于 channel 实现。



大多数的编程语言的并发编程模型是基于线程和内存同步访问控制，Go 的并发编程的模型则用 goroutine 和 channel 来替代。Goroutine 和线程类似，channel 和 mutex (用于内存同步访问控制)类似。

