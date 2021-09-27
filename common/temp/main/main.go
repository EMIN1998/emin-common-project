package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	//创建一个周期性的定时器
	ticker := time.NewTicker(3 * time.Second)

	//定义计数器
	count := 1
	fmt.Println("当前时间为:", time.Now(), "count = ", count)

	go func() {
		for {
			//从定时器中获取数据
			t := <-ticker.C
			count++
			fmt.Println("当前时间为:", t, "count = ", count)
			if count == 10 {
				/**
				  如果周期性定时被消费10次后就停止该定时器
				*/
				ticker.Stop()
				runtime.Goexit()
			}

		}
	}()


	time.Sleep(time.Second * 5)
	fmt.Println("main thread done!")
}
