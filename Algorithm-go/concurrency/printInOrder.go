package concurrency

type Action func(string)
type PrintInOrder interface {
	First(func(int) string)
	Second(func(int) string)
	Third(func(int) string)
}

type PrintInOrderGo struct {
	c1     chan bool
	c2     chan bool
	c3     chan bool
	Output string
}

func (p *PrintInOrderGo) Initial() {
	p.c1 = make(chan bool)
	p.c2 = make(chan bool)
	p.c3 = make(chan bool)
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
