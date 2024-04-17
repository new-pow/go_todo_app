package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	// "github.com/new-pow/go_todo_app/config"
	"go_todo_app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New() // 환경 변수를 읽어들인다.
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port)) // 포트를 지정하여 리스닝한다.
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux := NewMux()        // 라우팅 로직을 담은 mux를 생성한다.
	s := NewServer(l, mux) // 서버를 생성한다.
	return s.Run(ctx)      // 서버를 실행한다.
}
