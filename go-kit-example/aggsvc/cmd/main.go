package main

import (
	"net"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/olzh2102/toll-calculator/go-kit-example/aggsvc/aggendpoint"
	"github.com/olzh2102/toll-calculator/go-kit-example/aggsvc/aggservice"
	"github.com/olzh2102/toll-calculator/go-kit-example/aggsvc/aggtransport"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	service := aggservice.New(logger)
	endpoints := aggendpoint.New(service, logger)
	httpHandler := aggtransport.NewHTTPHandler(endpoints, logger)

	// The HTTP listener mounts the Go kit HTTP handler we created.
	httpListener, err := net.Listen("tcp", ":3000")
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "HTTP", "addr", ":3000")
	err = http.Serve(httpListener, httpHandler)
	if err != nil {
		panic(err)
	}
}
