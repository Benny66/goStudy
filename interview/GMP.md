```

```

- 编译成可执行文件，加载到虚拟地址空间中
- 检查、初始化
- runtime.main为执行入口，创建main.goroutine
- 调用main.main
- 全局变量：runtime.g runtime.m
- g0,m0,sched,allgs,allms,allps,runq
- exit()


goroutine的9种状态
- _Gidle=0:刚刚被分配，还没有被初始化。
- _Grunnable=1:在运行队列中
- _Grunning=2:可以执行用户代码
- _Gsyscall=3:正在执行一个系统调用
- _Gwaiting=4:在运行时被阻塞
- _Gmoribund_unused=5:当前未使用,但在 gdb 脚本中进行了硬编码
- _Gdead=6:当前未被使用,它可能刚刚退出，在空闲列表中，或者刚刚被初始化
- _Genqueue_unused=7:当前未使用
- _Gcopystack=8:堆栈正在被移动
- _Gpreempted=9:停止了自己的suspendG抢占,像 _Gwaiting，但还没有任何东西负责准备好它

分配到栈还是堆
- 不能确定分配对象的大小
- 对象生命周期超出当前所在函数


开始：白色，


