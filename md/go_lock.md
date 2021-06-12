# 锁
> 锁的目的就是为了并发编程中保证数据的安全

## 互斥锁 mutex
> mutual exclusion
> 必须同一时刻只能有一个 goroutine 对资源操作(读写)

```go
func main() {
    var lock *sync.Mutex
    
    lock.Lock()
    lock.Unlock()
}
```

## 读写锁
```go
func main() {
    var lock sync.RWMutex
    
    // 读锁
    lock.RLock()
    lock.RUnlock()
    
    // 写锁
    lock.Lock()
    lock.Unlock()
}
```

- 为了保证数据的安全，它规定了当有人还在读取数据（即读锁占用）时，不允计有人更新这个数据（即写锁会阻塞）
- 为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像 Mutex 那样只允许有一个人（线程）读取同一个数据。