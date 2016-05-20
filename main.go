package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/ogier/pflag"

	"github.com/denderello/ping-pong-http/api/v1"
	"github.com/denderello/ping-pong-http/server"
)

var (
	flags struct {
		host             string
		port             string
		shutdownCooldown time.Duration
		logLevel         string
	}
)

func init() {
	pflag.StringVar(&flags.host, "host", "", "Hostname to listen on")
	pflag.StringVar(&flags.port, "port", "9000", "Port to listen on")
	pflag.DurationVar(&flags.shutdownCooldown, "shutdown-cooldown", 10*time.Second, "Cooldown period to keep open connections alive before killing them")
	pflag.StringVar(&flags.logLevel, "log-level", "info", "Log level to use for log outputs")

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Ping Pong HTTP\n")
		pflag.PrintDefaults()
	}

	pflag.Parse()
}

func main() {
	l := logrus.StandardLogger()

	logLevel, err := logrus.ParseLevel(flags.logLevel)
	if err != nil {
		l.Fatal(err)
	}
	logrus.SetLevel(logLevel)

	s := server.New(server.ServerConfig{
		Addr: server.Address{
			Host: flags.host,
			Port: flags.port,
		},
		ShutdownCooldown: flags.shutdownCooldown,
		Logger:           l,
	})

	s.RegisterAPI(&v1.V1API{
		Logger: l,
	})

	sigs := make(chan os.Signal, 1)
	go func() {
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs

		s.Stop()
	}()

	s.Start()
}
