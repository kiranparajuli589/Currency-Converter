package server

import (
	"context"
	"fmt"
	protos "github.com/kiranparajuli589/currency-converter/protos/currency"
)

type Currency struct {
}

func NewCurrency() *Currency {
	return &Currency{}
}

func (c *Currency) GetRate(ctx context.Context, rateRequest *protos.RateRequest) (*protos.RateResponse, error) {
	fmt.Printf("Handle GetRate: 'base': %v, 'destination': %v",  rateRequest.GetBase(), rateRequest.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
