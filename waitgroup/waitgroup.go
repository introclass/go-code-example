package main

import (
	"sync"
)

//问题1：主函数先于 go routine 退出
//问题2：println(i) 打印的是 i 的当前值
//问题3：wg 被复制
func wg1() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(wg sync.WaitGroup) {
			wg.Add(1)
			println(i)
			wg.Done()
		}(wg)
	}
	wg.Wait()
}

func wg2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			println(i)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}

func wg3() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func wg4() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func wg5() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func wg6() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			println(i)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

func wg7() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg sync.WaitGroup) {
			println(i)
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
}

func main() {
	//wg1()
	//wg2()
	//wg3()
	//wg4()
	//wg5()
	//wg6()
	wg7()
}
