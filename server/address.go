package server

import "net"

type Addresser interface {
	Address() string
}

type Address struct {
	Host string
	Port string
}

func (a Address) Address() string {
	return net.JoinHostPort(a.Host, a.Port)
}
