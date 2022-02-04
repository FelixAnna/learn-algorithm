package concurrency

type PrintInOrderGo struct {
	c1     chan bool
	c2     chan bool
	c3     chan bool
	Output string
}

func (p *PrintInOrderGo) First(f func(int) string) {
	a := f(1)
	p.Output += a
	p.c1 <- true
}

func (p *PrintInOrderGo) Second(f func(int) string) {
	a := f(2)
	<-p.c1
	p.Output += a
	p.c2 <- true
}

func (p *PrintInOrderGo) Third(f func(int) string) {
	a := f(3)
	<-p.c2
	p.Output += a
	p.c3 <- true
}
