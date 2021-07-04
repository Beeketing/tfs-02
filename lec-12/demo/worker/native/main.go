package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"worker/mail"
	"worker/scheduler"
	"worker/worker"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

func main() {
	// prepare params
	apiKey := "" //TODO Change that!
	msgExchange := make(chan *mail.EmailContent)

	// prepare db
	db, err := sql.Open("mysql", "dev2:bdOtMZeXPyoT@tcp(127.0.0.1:3308)/dev") //TODO Change that!
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	// sql
	mailer := mail.NewSendgrid(apiKey)
	sched := scheduler.NewScheduler(ctx, db, msgExchange)
	worker := worker.NewWorker(ctx, wg, db, mailer, msgExchange)

	/////////////////////////////////////////////////
	// graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c // waits for the termination signal
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		sched.Stop() // stop scheduler at the end
		cancelFunc()
	}()
	/////////////////////////////////////////////////

	wg.Add(1) // add 1 for worker only. don't need for scheduler
	// run worker (as a receiver of msgExchange channel first)
	go worker.Start()

	sched.Start() // start to scan order

	// wait for the worker finishes its job
	wg.Wait()
}
