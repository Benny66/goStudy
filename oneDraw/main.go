package main

import (
    "fmt"
    "math/rand"
    "sort"
    "time"
)

// 参与者结构体
type Participant struct {
    Id        int       // 参与者ID
    Timestamp time.Time // 参与时间戳
    RandomNum int       // 随机数
}

func main() {
    // 初始化参与者列表
    participants := []Participant{
        {1, time.Now(), rand.Intn(100000)},
        {2, time.Now().Add(-time.Hour), rand.Intn(100000)},
        {3, time.Now().Add(-time.Minute), rand.Intn(100000)},
        {4, time.Now().Add(-time.Second), rand.Intn(100000)},
    }

    // 对参与者进行排序
    sort.Slice(participants, func(i, j int) bool {
        if participants[i].Timestamp.Equal(participants[j].Timestamp) {
            return participants[i].RandomNum < participants[j].RandomNum
        }
        return participants[i].Timestamp.Before(participants[j].Timestamp)
    })

    // 打印参与者顺序
    fmt.Println("参与者顺序：")
    for _, p := range participants {
        fmt.Printf("%v\n", p)
    }

    // 随机抽取中奖号码
    winnerIndex := rand.Intn(len(participants))
    winner := participants[winnerIndex]
    fmt.Printf("中奖号码：%v\n", winner)
}
