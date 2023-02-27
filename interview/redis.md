5种常见数据类型
- 字符串：共享 session、分布式锁，计数器、限流。
- 哈希：缓存用户信息
- 列表：栈，队列，消息队列，
- 集合：用户标签,生成随机数抽奖、社交需求
- 有序集合：排行榜，社交需求（如用户点赞）



为什么redis快呢？
- 内存实现，相比于MySQL的磁盘IO
- 不同数据类型适用不同的场景
- 合理的线程模型：IO多路复用，单线程不需要线程切换