# Interview

> 面向 **DevOps** | **SRE** | **DBA** 领域



## 1. 理论

### 1. 操作系统
- [tcp/udp/http](./md/tcp_udp_http.md)
- [IO多路复用](./md/io.md)
- [进程/线程/协程](./md/process_thread_coroutine.md)
- [http 状态码]
- [Cow技术]
- [同步/异步/阻塞/非阻塞](./md/sync_rsync_other.md)

### 2. Go

#### 初级

- [rune](./md/rune.md)
- [make 和 new 区别](./md/make_new.md)
- [defer](./md/defer.md)
- [func args](./md/func_value_point.md)
- [csp](./md/csp.md)
- [channel](./md/channel.md)
- [golang的锁](./md/go_lock.md)
- [单例模式](./md/singleton.md)
- [context](./md/context.md)
- [反射](./md/reflect.md)
- [go的命令行](./md/go_command.md)
- [比较两个接口类型是否相等](./md/interface_equal.md)
- [比较两个结构体是否相等](./md/struct_equal.md)

#### 中级

- golang 内存管理
- [逃逸分析](./md/escape.md)
- [go 为什么百万高并发](./md/goroutine.md)
- [gmp 调度模型](./md/gmp.md)
- [gc 垃圾回收](./md/go_gc.md)

#### 高级

- [go 性能调优](./md/go_debug.md)
- [reflect.DeepEqual 源码刨析](./md/source_code/reflect_deepequal.md)


### 3. Python

#### 初级

- [装饰器](./md/decorator.md)
- [iterator/ generator](./md/iter_gen.md)
- [with语法糖的实现方法]
- [元类]
- [Django 类里面mixin的用法]
- [谈谈 asyncio 理解]

#### 中级

- [GIL](./md/gil.md)
- [gc 垃圾回收](./md/python_gc.md)

#### 高级

- [python性能优化，如何提升python程序的运行效率](./md/python_performance.md)
- [python 代码是如何运行的](https://zhuanlan.zhihu.com/p/38855233)

### 4. 分布式
- [分布式一致性算法](./md/consistency_algorithm.md)
  - Raft
  - Gossip
  - Hash算法
  - Lease 租约机制
- [分布式系统如何保证数据一致性](./md/distributed.md)
- [CAP](./md/cap.md)  
- [etcd](./md/etcd.md)
- [consul](./md/consul.md)

### 5. 数据库
- [Redis](./md/redis.md)
- [MySQL](./md/mysql.md)
- [Kafka](./md/kafka.md)

### 6. 发布系统

- 发布策略
  - 滚动更新
  - 蓝绿发布
  - 金丝雀发布 / 灰度发布
- CI/CD
- 制品库





### 7. SRE

#### [1. 服务治理](./md/service_governance.md.md)

#### 2. 混沌工程

#### [3. Kubernetes](./md/kubernetes.md)



## 2. 数据结构和算法

### 1. 语法

- [冒泡](./data_structure/sort/bubble/main.go)
- [快排](./data_structure/sort/quick/main.go)
- 双端队列
- 链表
  - [单向链表](./data_structure/link_list/main.go)
  - [双向链表](./data_structure/double_link_list/main.go)
- 二叉树
  - [二叉树](./data_structure/binary_tree/main.go)
- [递归 && 分治](./data_structure/recursion/main.go)
- [二分查找](./data_structure/binary_search/main.go)


### 2. Leetcode

#### 1. 简单
- [1  两数之和](algorithm/two_sum/main.go)  
- [15 三数之和]
- [18 四数之和]  
- [11 盛最多水的容器]
- [20 有效的括号]
- [88 合并两个有序数组](algorithm/merge_array/main.go)
- [141 链表是否有环](algorithm/is_has_ring/main.go)
- [155 最小栈]
- [206 反转链表](algorithm/reverse_link/main.go)
- [234 回环链表](algorithm/is_palindrome/main.go)   
- [344 反转字符串](algorithm/reverse_string/main.go)  
- [509 斐波那契数列](algorithm/fibonacci/main.go)
- [617 合并二叉树] 
- [1143 最长公共子序列](algorithm/longest_common_subsequence/main.go)
- [283 移动零](algorithm/mv_zeroes/main.go)  
- [x 最大公约数](algorithm/gcd/main.go)
- [x 连续自然数差值](algorithm/different_one/main.go)  
- 26 删除有序数组中的重复项
#### 2. 中级
- [102 二叉树的层序遍历](algorithm/binary_tree_level_order_traversal/main.go)
- [107 二叉树的层序遍历 II](algorithm/binary_tree_level_order_traversal_2/main.go)  
- [goroutine 顺序打印1 到 9数字](algorithm/goroutine_order_print/main.go)
- [goroutine 交替打印字母和数字](algorithm/goroutine_alternate_print/main.go)
- [删除有序链表中重复的元素](algorithm/delete_link_duplicates/main.go)
- [字母排序](algorithm/sort_word/main.go)
- [交换单词](algorithm/swith_word/main.go)
- [遍历树的前三层](algorithm/for_level_tree/main.go)
- [nginx日志处理](algorithm/nginx/main.go)
- nginx负载均衡

#### 3. 高级

- [84 柱状图中最大的矩形]


#### 3. 高级
- [连续的自然数, 寻找丢失了一个数](./algorithm/find_num/main.go)



## 3. 书籍

### 1. 沟通管理

- 金字塔原理

### 2. 技术

- 代码整洁之道
- 干净架构
- Google SRE

### 3. 自我成长
- <跃迁>— 改变自己思维模式，提升认知，跳出单纯技术思维，去用更多的角度看待个人成长。
- <金字塔思维>—结构化思维，去分析问题，拆解问题。
- <非暴力沟通>—学会沟通，在沟通上应用方法论。
- <赋能>—打造应对不确定性敏捷团队，强调管理能力提升。
- <okr工作法>—系统的学习okr工作方法论。
- <刻意练习>—正确的、高效的学习方法论。
- <复盘>—复盘review的方法论
