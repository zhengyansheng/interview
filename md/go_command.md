# Go命令行



## 1. tool



### 1. 汇编代码

```bash
# go tool compile -S main.go


/*
	newobject 表示变量被分配到堆上

*/
```



### 2. 逃逸分析

```bash
✗ go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:12:13: ... argument does not escape
./main.go:12:13: b escapes to heap
```

