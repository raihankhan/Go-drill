package waitgroup

import "sync"

func ExecuteFunctionsInWaitGroup(fns ...func()) {
	var wg sync.WaitGroup

	for _, fn := range fns {
		wg.Add(1)
		go fn()
	}
	wg.Wait()
}

func ExecuteFunctionsNtimesInWaitGroup(fn func(), n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		fn()
	}
	wg.Wait()
}

func efd(f ...string) {
	println(f[0])
}
