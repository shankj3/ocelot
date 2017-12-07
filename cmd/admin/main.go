package main

import (
	"fmt"
	"github.com/namsral/flag"
	"github.com/shankj3/ocelot/util/cred"
	"github.com/shankj3/ocelot/admin"
	ocelog "bitbucket.org/level11consulting/go-til/log"
)

func main() {
	//load properties
	var port string
	var consulHost string
	var consulPort int
	var logLevel string

	flag.StringVar(&port, "port", "10000", "admin server port")
	flag.StringVar(&consulHost, "consul-host", "localhost", "consul host")
	flag.IntVar(&consulPort, "consul-port", 8500, "consul port")
	flag.StringVar(&logLevel, "log-level", "debug", "ocelot admin log level")
	flag.Parse()

	ocelog.InitializeLog(logLevel)

	serverRunsAt := fmt.Sprintf("localhost:%v", port)
	ocelog.Log().Debug(serverRunsAt)

	//TODO: this is my local vault root token, too lazy to set env variable
	configInstance, err := cred.GetInstance(consulHost, consulPort, "0175e9be-9917-a01f-cd65-b412f55ec545")

	if err != nil {
		ocelog.Log().Fatal("could not talk to consul or vault, bailing")
	}

	admin.Start(configInstance, serverRunsAt, port)

}
