package internal

import (
	"context"

	tseProto "github.com/2beens/tiny-stock-exchange-proto"
)

type TinyStockExchange struct {
	tseProto.UnimplementedTinyStockExchangeServer
}

func NewTinyStockExchange() *TinyStockExchange {
	return &TinyStockExchange{}
}

func (s *TinyStockExchange) NewStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) RemoveStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) UpdateStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) ListStocks(listStocksRequest *tseProto.ListStocksRequest, listStocksServer tseProto.TinyStockExchange_ListStocksServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) NewValueDelta(ctx context.Context, delta *tseProto.StockValueDelta) (*tseProto.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) ListStockValueDeltas(listStockValueDeltasRequest *tseProto.ListStockValueDeltasRequest, listStockValueDeltasServer tseProto.TinyStockExchange_ListStockValueDeltasServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) mustEmbedUnimplementedTinyStockExchangeServer() {
	//TODO implement me
	panic("implement me")
}
