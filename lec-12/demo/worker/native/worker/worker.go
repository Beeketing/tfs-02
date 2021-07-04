package worker

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"worker/mail"
)

// Worker defines a worker
type Worker struct {
	wg     *sync.WaitGroup
	mailer mail.Mailer
	inChan <-chan *mail.EmailContent
	ctx    context.Context
	db     *sql.DB
}

// NewWorker creates new worker
func NewWorker(ctx context.Context, wg *sync.WaitGroup, db *sql.DB, mailer mail.Mailer, ch <-chan *mail.EmailContent) *Worker {
	return &Worker{
		ctx:    ctx,
		wg:     wg,
		mailer: mailer,
		inChan: ch,
		db:     db,
	}
}

// Start starts worker to process message
//  Processing logic:
//    1. Wait for message
//    2. Send email with mailer (Sendgrid client)
//    3. Update database (thankyou_email_sent) to prevent duplicated emails
func (w *Worker) Start() {
	if w.mailer == nil || w.db == nil {
		fmt.Println("cannot start worker since mailer is nil")
		return
	}
	for {
		select {
		case em := <-w.inChan: // sending email now
			err := w.mailer.Send(em)
			if err != nil {
				fmt.Println("Cannot send email due to error: ", err)
				continue
			}
			// update sql data
			_, err = w.db.Exec("UPDATE `order` SET thankyou_email_sent = ? WHERE id = ?", true, em.ID)
			if err != nil {
				fmt.Println("Cannot update thankyou_email_sent to true")
			}
		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}
