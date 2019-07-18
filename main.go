package main

import (
	"finder/config"
	"finder/controllers"
	"finder/service"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	flagSet = flag.NewFlagSet("finder", flag.ExitOnError)

	filePath = flagSet.String("config-path", "config.toml", "config file")
)

func main() {
	flagSet.Parse(os.Args[1:])

	if err := config.Init(*filePath); err != nil {
		panic(err)
	}

	svr := service.New(config.Conf)
	//controller
	controllers.Init(svr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-c:
		svr.Close()
		return
	}
}
