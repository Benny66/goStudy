哈希
m(m个桶bucket)
hash(key)%m = 

哈希冲突，选择后面的桶

```
type hmap struct {
    count 键值对数量
    B   记录桶的数目，2^B

    oldbucktPtr     旧桶的地址
    bucket          当前桶的地址
    nevacuate       下一个迁移的旧桶编号

}


type bmap struct {
    
}
```