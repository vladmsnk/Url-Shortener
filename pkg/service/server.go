package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/url"
	"vladmsnk/urlshort/config"
	"vladmsnk/urlshort/internal/usecase"
	"vladmsnk/urlshort/pkg/grpcserver"
	"vladmsnk/urlshort/proto/pb"
)

type URLServer struct {
	pb.UnimplementedUrlShortenerServer
	uc *usecase.ShortenerUseCase
}

func NewURLShortenerServer(cfg *config.GRPc, uc *usecase.ShortenerUseCase) (*grpcserver.GRPCServer, error) {
	var opts []grpc.ServerOption

	grpcs, err := grpcserver.NewGRPCServer(cfg, opts...)
	if err != nil {
		return nil, err
	}
	pb.RegisterUrlShortenerServer(grpcs.Ser, URLServer{uc: uc})
	grpcs.Run()
	return grpcs, nil
}

func (s URLServer) CreateURL(context context.Context, createUrlRequest *pb.CreateURLRequest) (*pb.CreateURLResponse, error) {
	longURL := createUrlRequest.GetLongURL()

	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return &pb.CreateURLResponse{}, err
	}

	shortURL, err := s.uc.CreateURL(context, longURL)
	if err != nil {
		return &pb.CreateURLResponse{}, err
	}
	return &pb.CreateURLResponse{LongURL: longURL, ShortURL: shortURL}, nil
}

func (s URLServer) GetURL(context context.Context, getUrlRequest *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	shortURL := getUrlRequest.GetShortURL()

	longURL, err := s.uc.GetURL(context, shortURL)
	if err != nil {
		return &pb.GetURLResponse{}, err
	}

	header := metadata.Pairs("Location", longURL)
	err = grpc.SendHeader(context, header)
	if err != nil {
		return &pb.GetURLResponse{}, err
	}
	return &pb.GetURLResponse{ShortURL: shortURL, LongURL: longURL}, nil
}
