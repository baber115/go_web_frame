package demo

import (
	"fmt"
	"time"
)

// 轮询，每隔多少时间执行一次
/**
- 可以用ticker和sleep实现，但是ticker性能更好
- sleep是通过阻塞当前goroutine来实现的，需要先调度唤醒当前goroutine，才能继续后面的逻辑
- ticker在底层创建一个定时器，并且监听定时器产生的信号，继续后面的逻辑。
 - 如果有海量的定时器时，会进行统一调度，所以cpu消耗远小于sleep
*/
func tickerDemo() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for true {
		fmt.Println(1)
		<-ticker.C
	}
}
