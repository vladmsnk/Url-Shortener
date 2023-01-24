package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"vladmsnk/taskrec/config"
	"vladmsnk/taskrec/internal/usecase"
	"vladmsnk/taskrec/internal/usecase/repo"
	"vladmsnk/taskrec/pkg/grpcserver"
	"vladmsnk/taskrec/pkg/logger"
	"vladmsnk/taskrec/pkg/postgres"
	"vladmsnk/taskrec/pkg/service"
	"vladmsnk/taskrec/proto/pb"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	useCase := usecase.New(
		repo.New(pg),
	)

	go runRest(cfg.GRPc)
	grpcserver, err := service.NewURLShortenerServer(&cfg.GRPc, useCase)
	if err != nil {
		fmt.Println(err)
	}
	Lock(make(chan os.Signal, 1))
	grpcserver.Close()
}

func runRest(cfg config.GRPc) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(grpcserver.ResponseHeaderMatcher))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUrlShortenerHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), opts)
	if err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.RestPort), mux); err != nil {
		panic(err)
	}
}

func Lock(ch chan os.Signal) {
	defer func() {
		ch <- os.Interrupt
	}()
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-ch
}
