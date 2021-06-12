# make or new

- 参数不同
  - make -> slice, map, channel
  - new  -> int, string, struct  
- 返回值不同
  - make -> return Type
  - new -> return *Type

## make

```text
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.

func make(t Type, size ...IntegerType) Type

翻译
make 内置函数分配和初始化一个对象的类型仅 slice，map，channel，
和 new 相似的是 第一个参数是一个类型，不是值；不一样的是 make返回的是它参数一样的类型，不是它的指针
- slice
- map
- channel
```
## new

```text
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type

翻译
new 是内置函数分配内存，第一个参数是 类型，不是值
返回的是这个值的指针类型，分配该类型的零值
```