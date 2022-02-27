package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adityagoel/product-api/handlers"
)

func main() {
	thisLogger := log.New(os.Stdout, "product-api", log.LstdFlags)

	thisProductsHandler := handlers.NewProducts(thisLogger)

	thisServeMux := http.NewServeMux()

	thisServeMux.Handle("/", thisProductsHandler)

	thisServer := &http.Server{
		Addr:         ":8082",
		Handler:      thisServeMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		errorWhileStartingServer := thisServer.ListenAndServe()
		if errorWhileStartingServer != nil {
			thisLogger.Fatal(errorWhileStartingServer)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	thisSignalChannel := <-sigChan

	thisLogger.Println("Received Terminate, Graceful Shutdown", thisSignalChannel)

	timeOutContext, canFunct := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer canFunct() // releases resources if slowOperation completes before timeout elapses

	thisServer.Shutdown(timeOutContext)

}
