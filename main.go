package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SealTV/sovoiting/server"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// client, err := ethclient.Dial("http://127.0.0.1:8545")
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	nID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(nID.Int64())

	srv := http.Server{
		Addr:    ":8080",
		Handler: server.New().GetHttpHAndler(),
	}

	go func() {
		log.Println("Server started")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopCancel()

	if err := srv.Shutdown(stopCtx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
