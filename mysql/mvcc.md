MVCC，即Multi-Version  Concurrency Control （多版本并发控制）
参考：https://mp.weixin.qq.com/s?__biz=Mzg3NzU5NTIwNg==&mid=2247495277&idx=1&sn=a1812febb4246f824ce54d778f672025&chksm=cf223144f855b8528ad6cce707dc3a1b4d387817bd751dfab4f79dda90c6640f9763d25f3f33&token=1495321435&lang=zh_CN&scene=21#wechat_redirect


事务版本号：事务每次开启前，都会从数据库获得一个自增长的事务ID，可以从事务ID判断事务的执行先后顺序。这就是事务版本号。

隐藏字段：trx_id事务ID、roll_pointer指针，指向回滚段的undo日志、

快照读： 读取的是记录数据的可见版本（有旧的版本）。不加锁,普通的select语句都是快照读。
当前读：读取的是记录数据的最新版本，显式加锁的都是当前读。