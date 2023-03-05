package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"context"
	"syscall"
	"os"
	"os/signal"

	"pdbg.work/shiba/module/parse"
	
	"pdbg.work/shiba/module/setting"
	"pdbg.work/shiba/module/exec"
)

func Exec() {
	fmt.Println("start shiba system.")
	
	// Server.
	server := &http.Server{
		ReadHeaderTimeout: 7 * time.Second,
		Addr:    "localhost:8080",
		Handler: handleExec(parse.Parse),
	}
	
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	
	// Cron.
	var ch chan string
	var cronFlag int = 1

	if setting.TimerExec == true {
		ch = make(chan string)
		go func(ch <-chan string) {
			for {
				//
				select {
				case s := <-ch:			
					if s == "close" {
						break
					}
				default:
				}
				
				//
				if cronFlag == 1 {
					exec.Cron()
				} else {
					time.Sleep(time.Duration(setting.CronTimerMicroseccond) * time.Microsecond)
				}
				cronFlag = -cronFlag
			}
		}(ch)
	}
	
	// サーバーの終了処理	
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sigCh

	if setting.TimerExec == true {
		ch <- "close"
	}
	
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)	
	err := server.Shutdown(ctx) // Graceful Shutdown.
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Stop shiba system.")
	return
}

func handleExec(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        h(w, r)
    }
}