package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

// 값을 받아서 Server 구조체를 생성한다.
// 동적으로 포트를 선택하기 위해 net.Listener를 받는다.
// 라우팅 로직을 담은 mux를 받는다.
func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error { // 다른 고루틴에서 HTTP 서버를 실행한다.
		// http.ErrServerClosed 는 정상적인 종료를 의미한다.
		// http.Server.Shutdown() 이 호출되면 http.ErrServerClosed 가 반환되기 때문
		if err := s.srv.Serve(s.l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// 채널로부터의 알림(종료 알림)을 기다린다. 시그널이나 취소 요청에 의해 취소될 때까지 블록된 상태로 있다.
	<-ctx.Done()
	// 모든 요청을 처리한 후 서버를 종료한다.
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// 정상 종료를 기다린다.
	return eg.Wait()
}
