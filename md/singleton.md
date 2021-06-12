# 单例模式 singleton

## 特性
- 一个类只允许有一个实例子，节约资源

## 适用场景
- 配置文件
- 日志类


## sync.Once

```go

// 定义包私有结构体
type singleton struct {
}

var (
	once     *sync.Once
	instance *singleton
)

func GetInstance() *singleton {
	// 调用once.Do函数，进行实例化
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
```


## 例子

- 三个对象对地址不一样，空间浪费
```go
package main

import "fmt"

type WebConfig struct {
	Port int
}

func GetConfig() *WebConfig  {
	return &WebConfig{
		Port: 8000,
	}
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()
	c3 := GetConfig()
	fmt.Println(c1 == c2, c2 == c3) // 三个对象对地址不一样
}
```

- 手写单例子模式
```go
package main

import "fmt"

type WebConfig struct {
	Port int
}

var webConfig *WebConfig

func GetConfig() *WebConfig {
	if webConfig == nil {
		webConfig = &WebConfig{
			Port: 8000,
		}
		return webConfig
	}
	return webConfig
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()
	c3 := GetConfig()
	fmt.Println(c1 == c2, c2 == c3)
}

```

```go
package main

import (
	"fmt"
	"sync"
)

type WebConfig struct {
	Port int
}

var webConfig *WebConfig
var lock sync.Mutex

func GetConfig() *WebConfig {
	lock.Lock()
	defer lock.Unlock()
	if webConfig == nil {
		webConfig = &WebConfig{
			Port: 8000,
		}
		return webConfig
	}
	return webConfig
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()
	c3 := GetConfig()
	fmt.Println(c1 == c2, c2 == c3)
}

```