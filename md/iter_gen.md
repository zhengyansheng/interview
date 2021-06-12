## iterator
> 迭代器  

如果一个对象通过 dir 内置方法查看，具有 __iter__ 和 __next__ 方法，就称这个对象是迭代器；

### 特性

- 通过 next 方法可以查看对象的下一个对象
- 如果这个对象已经是最后一个对象，最后会抛 StopIteration 异常

```bash
>>> dir(arr)
__class__
__delattr__
__dir__
__doc__
__eq__
__format__
__ge__
__getattribute__
__gt__
__hash__
__init__
__init_subclass__
__iter__   # 方法1
__le__
__length_hint__
__lt__
__ne__
__new__
__next__  # 方法2
__reduce__
__reduce_ex__
__repr__
__setattr__
__setstate__
__sizeof__
__str__
__subclasshook__
````
## generator

生成器
如果有 yield 关键字，就成为生成器

### 特性
- 节省内存，按需读取，不用一次性加载全部到内存；