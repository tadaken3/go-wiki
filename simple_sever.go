package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func routes() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world.")
	})

	return
}

func server(addr string) (listener net.Listener, ch chan error) {
	ch = make(chan error)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		mux := routes()
		ch <- http.Serve(listener, mux)
	}()

	return
}

func main() {
	listener, ch := server(":8080")
	fmt.Println("Server started at", listener.Addr())

	// シグナルハンドリング (Ctrl + C)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		log.Println(<-sig)
		listener.Close()
	}()

	log.Println(<-ch)
}
