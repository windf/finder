package main

import (
	"finder/config"
	"finder/service"
	"flag"
	"os"
)

var (
	flagSet = flag.NewFlagSet("finder", flag.ExitOnError)

	filePath    = flagSet.String("config-path", "config.toml", "config file")
	LogLevel    = flagSet.String("log-level", "DEBUG", "log level")
	httpAddress = flagSet.String("http-address", "0.0.0.0:8080", "<addr>:<port> to listen on for http clients")
)

func main() {
	flagSet.Parse(os.Args[1:])

	if err := config.Init(*filePath); err != nil {
		panic(err)
	}

	if err := service.Run(*httpAddress); err != nil {
		panic(err)
	}
}
