# Python gc 垃圾回收
> garbage collector

## 引用计数 为主

### 优点
实现简单

### 缺点
无法解决对象间的循环引用问题

## 标记和清除 为辅

- 标记阶段
- 清除阶段

### 分代回收 为辅
> Generational Collection

对象按照生命周期的长短可分三代
- 第一代 
- 第二代 
- 第三代 

```pythonregexp
>>> import gc
>>> gc.get_threshold() # 三代的阈值分别是700，10，10
(700, 10, 10)
>>>
>>> gc.get_count()     # 当前三代中对象的个数是673，8，0 只有达到阈值时才会触发 gc回收
(673, 8, 0)

>>> gc.collect()       # 手动触发gc回收
0
>>> gc.get_count()    
(4, 0, 0)
```


## 总结
Python的垃圾回收主要通过 引用计数 为主，标记和清除为辅解决对象循环间的引用，通过 分代回收 以空间换时间提高垃圾回收效率。

## 参考
- [Python垃圾回收机制](https://zhuanlan.zhihu.com/p/83251959)
- [Python垃圾收集：它是什么以及如何工作](https://stackify.com/python-garbage-collection/)