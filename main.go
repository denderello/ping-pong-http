package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/denderello/ping-pong-http/api/v1"
	"github.com/denderello/ping-pong-http/server"
)

func main() {
	s := server.New(server.ServerConfig{
		Addr: server.Address{
			Port: "9000",
		},
		ShutdownCooldown: 10 * time.Second,
		Logger:           logrus.StandardLogger(),
	})

	s.RegisterAPI(v1.V1API{})

	sigs := make(chan os.Signal, 1)
	go func() {
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs

		s.Stop()
	}()

	s.Start()
}
