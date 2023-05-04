package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     true,
	}
}

func withTLS(opt *Opts) {
	opt.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func withID(s string) OptFunc {
	return func(opts *Opts) {
		opts.id = s
	}
}

type Server struct {
	Opts
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	fmt.Printf("default: %+v\n", NewServer())
	custom := NewServer(withTLS, withID("custom"), withMaxConn(99))
	fmt.Printf("custom: %+v\n", custom)
}
