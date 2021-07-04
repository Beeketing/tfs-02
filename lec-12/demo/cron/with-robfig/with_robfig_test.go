package robfigjob

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func init() {
	c = cron.New(cron.WithSeconds())
}

func TestRobfigSimpleJob(t *testing.T) {
	defer c.Stop()
	c.AddFunc("* * * * * *", doSimplejob)
	c.Start()
	time.Sleep(time.Second * 5)
	t.Error("done")
}

func TestRobfigTimeConsumedJob(t *testing.T) {
	defer c.Stop()
	c.AddFunc("* * * * * *", doTimeConsumedJob)
	c.Start()
	time.Sleep(time.Second * 10)
	t.Error("done")
}

func doSimplejob() {
	fmt.Printf("I'm doing my simple work at: %v\n", getCurrentString())
}

func doTimeConsumedJob() {
	fmt.Printf("I'm doing my time consumed work at: %v\n", getCurrentString())
	// do
	time.Sleep(time.Second * 3)
}

func getCurrentString() string {
	return time.Now().Format("2006-Jan-02 15:04:05")
}
