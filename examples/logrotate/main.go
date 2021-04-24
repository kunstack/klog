package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/kunstack/klog"
)

func main() {
	log.SetLevel(log.DebugLevel) // set log level to debug
	defer log.Flush()            // flush all buffer before exit

	file, err := os.OpenFile("./logrotate.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	// defer file.Close()
	log.SetOutput(file) //Initialize the log file

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	rotateHandler := make(chan os.Signal)
	signal.Notify(rotateHandler, syscall.SIGUSR1) // log rotate signal,  kill -SIGUSR1 $pid

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-rotateHandler:
				log.SetOutput(os.Stderr) // Temporarily set to os.Stderr
				if err := file.Close(); err != nil {
					log.Errorln(err)
					break
				}
				file, err = os.OpenFile("./logrotate.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					log.Errorln(err)
					break
				}
				log.SetOutput(file)
				log.Debugln("Log rotate was successful")
			case <-ctx.Done(): //Receive stop signal e.g. ctl+c
				close(rotateHandler)
				return
			}
		}
	}()

	wg.Add(1)
	go func() { // output log
		defer wg.Done()
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			log.Debugf("this is This is the %dth cycle", i)
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()

	wg.Wait()
}
