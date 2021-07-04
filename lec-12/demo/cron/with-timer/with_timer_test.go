package timerjob

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleJob(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("exiting")
				return
			case <-ticker.C:
				doSimplejob()
			}
		}
	}()

	time.Sleep(time.Second * 5)
	done <- true
	t.Error(1) // for printing previous logs
}

func TestTimeConsumedJob(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("exiting")
				return
			case <-ticker.C:
				fmt.Printf("Started at: %v\n", getCurrentString())
				doTimeConsumedJob()
				doSimplejob()
				fmt.Printf("Finish at: %v\n", getCurrentString())
			}
		}
	}()

	time.Sleep(time.Second * 10)
	done <- true
	t.Error(1) // for printing previous logs
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
