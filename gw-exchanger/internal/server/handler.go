package server

import (
	"context"

	pb "github.com/chtozamm/javacode-final/proto-exchange/exchange"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetExchangeRates(ctx context.Context, req *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.db.GetAllExchangeRates()
	if err != nil {
		s.log.Err(err).Msg("Failed to get exchange rates")
		return nil, status.Errorf(codes.Internal, "failed to get exchange rates: %v", err)

	}

	ratesResponse := &pb.ExchangeRatesResponse{
		Rates: rates,
	}

	return ratesResponse, nil
}

func (s *server) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	rate, err := s.db.GetExchangeRateForCurrency(req.FromCurrency, req.ToCurrency)
	if err != nil {
		s.log.Err(err).Msg("Failed to get exchange rates")
		return nil, status.Errorf(codes.Internal, "failed to get exchange rates: %v", err)
	}

	rateResponse := &pb.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         rate,
	}

	return rateResponse, nil
}
