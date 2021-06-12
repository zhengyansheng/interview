# GIL
> Global Interpreter Lock 
> 全局解释器锁

## GIL特性
- 互斥锁
- 保证同一个进程中只允许有1个线程执行，目的是解决多线程中操作全局变量引起的数据不安全问题

## 常见的解释器
- CPython: C 语言到Python解释器      
  - 线程切换
    - 网络IO堵塞
    - 抢占式多任务处理能力，每1000个字节码会中断一个线程。
- JPython: Java 语言到Python解释器
- IronPython: C# 语言到Python解释器

## Parallelism
GIL 禁止并行任务

## 非原子操作
```python
import dis

n = 1

def foo():
    global n
    n += 1

dis.dis(foo)
```

```text
  8           0 LOAD_GLOBAL              0 (n)   # 加载全局变量n放到栈上
              2 LOAD_CONST               1 (1)   # 加载常量1放到栈上
              4 INPLACE_ADD                      # 栈上最顶层的2个变量相加
              6 STORE_GLOBAL             0 (n)   # 相加的和在存储到全局变量n上
              8 LOAD_CONST               0 (None)
             10 RETURN_VALUE
```

通过内置函数 dis 来把Python代码翻译成字节码时，可以看到 n+=1 并非是原子操作而是分成了四步来完成

## 参考
- [什么是Python全局解释器锁（GIL）](https://realpython.com/python-gil/#:~:text=The%20Python%20Global%20Interpreter%20Lock%20or%20GIL%2C%20in%20simple%20words,at%20any%20point%20in%20time.)
- [Grok the GIL：如何编写快速且线程安全的Python](https://opensource.com/article/17/4/grok-gil)
- [谈谈python的GIL、多线程、多进程](https://zhuanlan.zhihu.com/p/20953544)