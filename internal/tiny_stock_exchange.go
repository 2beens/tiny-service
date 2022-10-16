package internal

import (
	"context"
)

type tinyStockExchangeServer struct {
}

func newTinyStockExchangeServer() *tinyStockExchangeServer {

}

func (s *tinyStockExchangeServer) NewStock(context.Context, *Stock) (*Result, error) {

}

func (s *tinyStockExchangeServer) RemoveStock(context.Context, *Stock) (*Result, error) {

}

func (s *tinyStockExchangeServer) UpdateStock(context.Context, *Stock) (*Result, error) {

}

func (s *tinyStockExchangeServer) ListStocks(*ListStocksRequest, TinyStockExchange_ListStocksServer) error {

}

func (s *tinyStockExchangeServer) NewValueDelta(context.Context, *StockValueDelta) (*Result, error) {

}

func (s *tinyStockExchangeServer) ListStockValueDeltas(*ListStockValueDeltasRequest, TinyStockExchange_ListStockValueDeltasServer) error {

}
