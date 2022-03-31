package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/pflag"
	"github.com/zyyw/hello-expo/log"
	"github.com/zyyw/hello-expo/pkg"
	"github.com/zyyw/hello-expo/routers"
)

var (
	cfg = pflag.StringP("config", "c", "", "hello-expo config file path")
)

// GitCommit holds the git commit hash used in the build
var GitCommit string

func main() {
	pflag.Parse()

	// init config via viper
	if err := pkg.InitConfig(*cfg); err != nil {
		fmt.Printf("hello-expo version: %s\n", GitCommit)
		panic(err)
	}

	// Print the Git Hash.
	log.Logger().Sugar().Infof("hello-expo version: %s", GitCommit)

	router := routers.InitRouter()

	log.Logger().Sugar().Infof("starting a server listening on port %d", pkg.Conf.Server.HttpPort)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", pkg.Conf.Server.HttpPort),
		Handler: router,
		//TLSConfig:         nil,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}

	server.ListenAndServe()
}
