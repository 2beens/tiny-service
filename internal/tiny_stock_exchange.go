package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	tseProto "github.com/2beens/tiny-stock-exchange-proto"

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

	if err := CreateIndex(collStocks, "ticker", true); err != nil {
		return nil, fmt.Errorf("create index for stock.ticker: %w", err)
	}

	return &TinyStockExchange{
		mongoClient:     mongoClient,
		collStocks:      collStocks,
		collValueDeltas: collValueDeltas,
	}, nil
}

// CreateIndex - creates an index for a specific field in a collection
// Creating indexes in MongoDB is an idempotent operation, so no need to check if it exists:
// https://stackoverflow.com/a/35020346/1794478
func CreateIndex(collection *mongo.Collection, field string, unique bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if createdIndex, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		// index in ascending order or -1 for descending order
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s for field %s", createdIndex, field)
	}

	return nil
}

func (s *TinyStockExchange) NewStock(ctx context.Context, stock *tseProto.Stock) (*tseProto.Result, error) {
	log.Printf("> will try to insert stock: %v", stock)
	result, err := s.collStocks.InsertOne(ctx, bson.D{
		{"ticker", stock.Ticker},
		{"name", stock.Name},
	})
	if err != nil {
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
		return nil, err
	}

	msg := fmt.Sprintf("deleted documents (stock: %s): %v", stock.Ticker, result.DeletedCount)
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
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) NewValueDelta(ctx context.Context, delta *tseProto.StockValueDelta) (*tseProto.Result, error) {
	log.Printf("> will try to insert value delta: %v", delta)
	result, err := s.collValueDeltas.InsertOne(ctx, bson.D{
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
	//TODO implement me
	panic("implement me")
}

func (s *TinyStockExchange) mustEmbedUnimplementedTinyStockExchangeServer() {
	//TODO implement me
	panic("implement me")
}
