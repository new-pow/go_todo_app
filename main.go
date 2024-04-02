package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	// 동작 확인을 위해
	"golang.org/x/sync/errgroup"

	"github.com/new-pow/go_todo_app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM) // os.Interrupt 대신 syscall.SIGINT 사용
	// defer stop()                                                            // defer로 시그널이 들어오면 stop
	cfg, err := config.New() // config 생성
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port)) // 포트 설정
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("server listening at %s", url)
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 명령 줄에서 테스트 하기 위한 로직
			time.Sleep(5 * time.Second)
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	// 다른 고루틴에서 HTTP 서버를 실행한다.
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})
	// 채널로부터의 알림(종료 알림)을 기다린다.
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Go 메서드로 실행한 다른 고루틴의 종료를 기다린다.
	return eg.Wait()
}
