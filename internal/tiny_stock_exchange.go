package internal

import (
	"context"
	"fmt"

	"github.com/2beens/tiny-service/pkg"
	tseProto "github.com/2beens/tiny-stock-exchange-proto"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TinyStockExchange struct {
	tseProto.UnimplementedTinyStockExchangeServer

	mongoClient     *mongo.Client
	collStocks      *mongo.Collection
	collValueDeltas *mongo.Collection
}

func NewTinyStockExchange(
	dbName string,
	mongoClient *mongo.Client,
) (*TinyStockExchange, error) {
	tseDB := mongoClient.Database(dbName)
	collStocks := tseDB.Collection("stocks")
	collValueDeltas := tseDB.Collection("deltas")

	if err := pkg.CreateIndex(collStocks, "ticker", true); err != nil {
		return nil, fmt.Errorf("create index for stock.ticker: %w", err)
	}

	return &TinyStockExchange{
		mongoClient:     mongoClient,
		collStocks:      collStocks,
		collValueDeltas: collValueDeltas,
	}, nil
}

func (s *TinyStockExchange) NewStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	log.Printf("> will try to insert stock: %v", stock)
	result, err := s.collStocks.InsertOne(ctx, bson.D{
		{"ticker", stock.Ticker},
		{"name", stock.Name},
	})
	if err != nil {
		log.Printf("add new stock %s: %s", stock.Ticker, err)
		return nil, err
	}

	msg := fmt.Sprintf("inserted document (stock) with _id: %v", result.InsertedID)
	fmt.Println(msg)

	return &tseProto.Result{
		Success: true,
		Message: msg,
	}, nil
}

func (s *TinyStockExchange) RemoveStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	log.Printf("> will try to remove stock: %v", stock)
	filter := bson.D{
		{"ticker", bson.D{{"$eq", stock.Ticker}}},
	}
	opts := options.Delete().SetHint(bson.D{{"ticker", 1}})

	result, err := s.collStocks.DeleteOne(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to delete stock %s: %w", stock.Ticker, err)
	}
	deletedStocksCount := result.DeletedCount

	result, err = s.collValueDeltas.DeleteMany(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to delete related delta values for stock %s: %w", stock.Ticker, err)
	}
	deletedValueDeltas := result.DeletedCount

	msg := fmt.Sprintf("deleted documents (stock: %s) %d, value deltas %d", stock.Ticker, deletedStocksCount, deletedValueDeltas)
	fmt.Println(msg)

	return &tseProto.Result{
		Success: true,
		Message: msg,
	}, nil
}

func (s *TinyStockExchange) UpdateStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) ListStocks(listStocksRequest *tseProto.ListStocksRequest, listStocksServer tseProto.TinyStockExchange_ListStocksServer) error {
	filter := bson.D{}
	cursor, err := s.collStocks.Find(listStocksServer.Context(), filter)
	if err != nil {
		return err
	}

	var results []bson.D
	if err := cursor.All(listStocksServer.Context(), &results); err != nil {
		return err
	}

	for _, result := range results {
		resMap := result.Map()
		listStocksServer.Send(&tseProto.Stock{
			Ticker: fmt.Sprintf("%s", resMap["ticker"]),
			Name:   fmt.Sprintf("%s", resMap["name"]),
		})
	}

	return nil
}

func (s *TinyStockExchange) NewValueDelta(ctx context.Context, delta *tseProto.StockValueDelta) (*tseProto.Result, error) {
	log.Printf("> will try to insert value delta: %v", delta)
	result, err := s.collValueDeltas.InsertOne(ctx, bson.D{
		{"ticker", delta.Ticker},
		{"delta", delta.Delta},
		{"timestamp", delta.Timestamp},
	})
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("inserted document (value delta) with _id: %v", result.InsertedID)
	fmt.Println(msg)

	return &tseProto.Result{
		Success: true,
		Message: msg,
	}, nil
}

func (s *TinyStockExchange) ListStockValueDeltas(listStockValueDeltasRequest *tseProto.ListStockValueDeltasRequest, listStockValueDeltasServer tseProto.TinyStockExchange_ListStockValueDeltasServer) error {
	filter := bson.D{}
	cursor, err := s.collValueDeltas.Find(listStockValueDeltasServer.Context(), filter)
	if err != nil {
		return err
	}

	var results []bson.D
	if err := cursor.All(listStockValueDeltasServer.Context(), &results); err != nil {
		return err
	}

	for _, result := range results {
		resMap := result.Map()
		ts, ok := resMap["timestamp"].(int64)
		if !ok {
			log.Errorf("list stock deltas: timestamp [%v] not an int64!", resMap["timestamp"])
			ts = 0 // or maybe skip
		}
		delta, ok := resMap[""].(int64)
		if !ok {
			log.Errorf("list stock deltas: delta [%v] not an int64!", resMap["delta"])
			delta = 0 // or maybe skip
		}
		listStockValueDeltasServer.Send(&tseProto.StockValueDelta{
			Ticker:    fmt.Sprintf("%s", resMap["ticker"]),
			Timestamp: ts,
			Delta:     delta,
		})
	}

	return nil
}
