package unixpipe

import "fmt"
import "sync"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func Exec() {
	nums := gen(1, 2, 3)
	out := sq(nums)

	fmt.Println(<-out)
	fmt.Println(<-out)
}

func Exec2() {
	in := gen(1, 2, 3)

	// 启动两个sq实例,即两个goroutines处理channel in的数据
	c1 := sq(in)
	c2 := sq(in)

	// merge函数将channel c1 和 c2合并在一起,这段代码会消费merge的结果
	for n := range merge(c1, c2) {
		fmt.Println("num from merge return: ", n)
	}
}

func Exec3() {
	done := make(chan struct{})
	defer close(done)
	in := gen1(done, 2, 3)

	c1 := sq1(done, in)
	c2 := sql(done, in)

}

func gen1() {

}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen2() {

}
