package main

//go:generate mockgen -source=main.go -destination=./my_pckg/my_pckg_mock.go -package=my_pckg

func main() {
}

type Simpler interface {
	SimpleMethod() // this will be run within main goroutine
	ConcurrentMethod()
}

func NeedTest(ms Simpler) Simpler {
	ms.SimpleMethod()
	go func(ms Simpler) {
		ms.ConcurrentMethod()
	}(ms)
	// time.Sleep(1 * time.Second) this will fix "test" but it's too artificial
	return nil
}
