package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Printf("🌋: Hello World!\n💻: %s\n📂: %s\n⏰: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))

	ctx, cancel := context.WithCancel(context.Background())

	var (
		data   int
		readCh = make(chan struct{})
	)
	go func() {
		log.Println("blocked to read data")
		// fake long i/o operations
		time.Sleep(3 * time.Second)
		data = 10
		log.Println("done read data")

		readCh <- struct{}{}
	}()

	time.AfterFunc(5*time.Second, cancel)

	select {
	case <-ctx.Done():
		log.Println("cancelled")
		return
	case <-readCh:
		break
	}

	log.Println("got final data", data)
}

/*

	times++
	var (
		data   responseData
		err    error
		readCh = make(chan struct{})
	)

	// read or cancel
	go func() {
		err = c.conn.ReadJSON(&data)
		readCh <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		c.log.Debugw("listener is cancelled", zap.Error(ctx.Err()))
		break LOOP
	case <-readCh:
		break
	}

	l := c.log.With("times", times, "past_time", time.Since(startTime), "interval_time", time.Since(lastTime), "message", data)
	lastTime = time.Now()

	// handle read error
	if err != nil {
		if ws.IsCloseError(err, 1006) {
			l.Warnw("websocket connection closed", "times", times, zap.Error(err))
		} else {
			l.Warnw("fail to read message", "times", times, zap.Error(err))
		}
		break LOOP
	}


*/
