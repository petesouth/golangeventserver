package truevieweventserver

import "os"
import "os/signal"
import "syscall"
import "log"

func WaitForSystemSignals() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				log.Println("hungup")
				exit_chan <- 1

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				log.Println("CTRL+C Goodbye")
				exit_chan <- 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				log.Println("force stop Goodebye")
				exit_chan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				log.Println("stop and core dump Goodbye")
				exit_chan <- 0

			default:
				log.Println("Unknown signal. Goodbye")
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)
}
