package server

import (
	"context"

	pb "github.com/chtozamm/javacode-final/proto-exchange/exchange"
)

type Server struct {
	pb.UnimplementedExchangeServiceServer
}

func (s *Server) GetExchangeRates(ctx context.Context, req *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	// TODO: retrieve exchange rates from database
	rates := map[string]float32{
		"USD": 1,
		"EUR": 0.92008,
		"RUB": 85.84,
	}

	ratesResponse := &pb.ExchangeRatesResponse{
		Rates: rates,
	}

	// TODO: handle potential errors using status package
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "")
	// }

	return ratesResponse, nil
}

func (s *Server) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	// TODO: retrieve exchange rate from database
	rates := map[string]float32{
		"USD": 1,
		"EUR": 0.92008,
		"RUB": 85.84,
	}
	rate := rates[req.ToCurrency] / rates[req.FromCurrency]

	rateResponse := &pb.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         rate,
	}

	return rateResponse, nil
}
