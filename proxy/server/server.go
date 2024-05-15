package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"proxy/config"

	"go.uber.org/zap"
)

type HTTPServer struct {
	conf   config.Server
	srv    *http.Server
	logger *zap.Logger
}

func NewHttpServer(conf config.Server, server *http.Server, logger *zap.Logger) *HTTPServer {
	return &HTTPServer{conf: conf, srv: server, logger: logger}
}

func (s *HTTPServer) Serve() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		s.logger.Info("server started", zap.String("port", s.conf.Port))
		listener, err := net.Listen("tcp", s.srv.Addr)
		if err != nil {
			s.logger.Fatal("failed to listen: %s", zap.Error(err))
			return
		}
		s.srv.Serve(listener)
	}()

	// Отправить сигнал в контейнер -> docker stop --help
	sig := <-quit
	s.logger.Info("received signal:", zap.Any("signal", sig))

	// Создание контекста с таймаутом для graceful shutdown
	ctx, shutdown := context.WithTimeout(context.Background(), s.conf.ShutdownTimeout)
	defer shutdown()

	// Остановка сервера с использованием graceful shutdown
	s.srv.Shutdown(ctx)
	s.logger.Info("server stopped gracefully")
}
