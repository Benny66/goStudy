### mysql

SQL 优化
- 索引，针对一些业务查询比较慢的功能，对sql进行慢查询日志分析（explain），判断是否已使用索引，使用索引是否失效了，在索引这个方向上做sql优化
- 尽量不在事务中做查询操作，尽可能在事务前将数据查询出来，如果查询时间过长会导致事务持续时间长，影响业务进行
- select查询一条数据的时候使用limit 1
- 如果使用表关联join查询，注意关联字段使用索引，两个表的字符集保持一致，不然索引失效；小表驱动大表
- 注意where查询的范围，in查询数据不应该过多，条件中使用函数会索引失效
- update数据的时候添加order by排序
- insert尽量选择批量SQL插入，数据较多可多次提交事务

针对 SQL 语句如何建立索引（最左原则）
- 外键的字段一定要建立索引
- where经常查询的字段建立索引
- order by、group by、distinct的字段要建立索引
- 索引不包含有NULL值的列

索引数据结构（B+树）(innodb)
![B+树](https://img-blog.csdnimg.cn/img_convert/9b7dc733fa94a2f1db7956baf0b9ade7.png)
- B+树非叶子节点上是不存储数据的，仅存储键值
- B+树所有数据均存储在叶子节点

- 以主键作为B+树索引的键值就是聚集索引
- 以主键以外的列值作为键值构建的B+树索引就是非聚集索引，叶子节点上不保存数据，保存对应主键的ID


### redis

redis 的使用场景？
- 做数据缓存：比如用户信息
- 分布式锁：微信支付成功回调消息，给订单唯一标识订单号加锁
- 计数器：社区文章的点赞、浏览量，允许一定的延迟，读入redis再写入数据库
- - string、hash和sorted set都提供了incr方法用于原子性的自增操作
- 消息队列

内存不足有新数据缓存，redis 如何处理？
- Redis会检查内存使用，如果内存使用超过 maxmemory，就会按照置换策略删除一些 key
新的命令执行成功

redis 的淘汰策略
- lru 最近很少的使用的key（根据时间，最不常用的淘汰）
- - 随机采集淘汰的key，每次随机选出5个key（默认），然后淘汰这5个key中最少使用的key
- lfu 最近很少的使用的key (根据计数器，用的次数最少的key淘汰)
- random 随机淘汰
- ttl 快要过期的先淘汰

redis 的缓存一致性
- 事务机制，保证更新缓存和更新数据库是原子性的。
- 线程A删除缓存后未来得及更新数据，被线程B查询数据并写入缓存，使数据库数据和缓存数据不一致，解决办法是双删，延迟双删，线程A延迟时间大于线程B查询并写入缓存时间

### go

线上服务内存泄漏如何处理？

go 为啥使用 CSP 模型来实现并发？

channel 是线程安全的吗？

channel 怎么保证的线程安全？

channel 的底层数据结构

mutex 是悲观锁还是乐观锁？

rwmutex 的使用场景

map 是线程安全的吗

无锁，但是线程安全的 map 如何设计？（CAS）

结构体能否进行比较？

空结构体的使用场景

给 channel 用空结构体的好处是什么？

字符串转 byte 数组会发生内存拷贝吗？

字符串转 byte 数组，如何避免发送内存拷贝？

多核 CPU 如何保持 cache 一致？（MESI协议）

GMP 调度模型，GMP 中的抢占机制

GC 三色标记，stw 是什么意思？产生 stw 的原因是什么？小对象多了，为什么会造成 GC 压力？

#### 代码题

go 函数中，返回值未命名，发生了 panic，但是在函数内 recover 了，函数返回什么值？怎么解决？

```

func test() error { var err error defer func() { if r := recover(); r != nil { err = errors.New(fmt.Sprintf("%s", r)) } }() raisePanic() return err } func raisePanic() { panic("发生了错误") } 

```

打印值是多少？

```

func main() { var a uint = 0 var b uint = 1 c := a - b fmt.Println(reflect.TypeOf(c)) fmt.Println(c) } 

```

```

func main() { a := [3]int{1, 2, 3} // 数组 for k, v := range a { if k == 0 { a[0], a[1] = 100, 200 fmt.Print(a) } a[k] = 100 + v } fmt.Print(a) } 

```

```

func main() { a := []int{1, 2, 3} // 切片 for k, v := range a { if k == 0 { a[0], a[1] = 100, 200 fmt.Print(a) } a[k] = 100 + v } fmt.Print(a) } 

```

每个函数起一个 goroutine，轮流打印 cat、dog、fish 100 次

手写 LRU 算法

网络编程

tcp 三次握手，四次挥手，四次挥手中 time_wait 可怕还是 close_wait 可怕

用过哪些 Linux 命令

tcp 的流量控制

rpc 是什么？如何设计一个 rpc？

如何处理 tcp 粘包问题？

http1.1 和 http2 的区别