package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/kunstack/klog"
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	stopChan := make(chan struct{})

	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		close(stopChan)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stopChan
}

func main() {
	log.SetLevel(log.DebugLevel) // set log level to debug
	defer log.Flush()            // flush all buffer before exit

	file, err := os.OpenFile("./logrotate.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file) //Initialize the log file

	stopChan := setupSignalHandler()

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
			case <-stopChan: //Receive stop signal e.g. ctl+c
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
			case <-stopChan:
				return
			default:
			}
		}
	}()

	wg.Wait()
}
