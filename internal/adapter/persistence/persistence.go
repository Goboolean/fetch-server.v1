package persistence

import (
	"context"

	"github.com/Goboolean/fetch-server/internal/domain/entity"
	"github.com/Goboolean/fetch-server/internal/domain/port"
	"github.com/Goboolean/fetch-server/internal/domain/port/out"
	"github.com/Goboolean/shared/pkg/mongo"
	"github.com/Goboolean/shared/pkg/rdbms"
)




type Adapter struct {
	rdbms *rdbms.Queries
	mongo *mongo.Queries
}

func NewAdapter(rdbms *rdbms.Queries, mongo *mongo.Queries) out.StockPersistencePort {
	return &Adapter{
		rdbms: rdbms,
		mongo: mongo,
	}
}



func (a *Adapter) StoreStock(tx port.Transactioner, stockId string, agg *entity.StockAggregate) error {

	dto := &mongo.StockAggregate{
		EventType: agg.EventType,
		Avg:       agg.Average,
		Min:       agg.Min,
		Max:       agg.Max,
		Start:     agg.Start,
		End:       agg.End,
		StartTime: agg.StartTime,
		EndTime:   agg.EndTime,
	}
	
	return a.mongo.InsertStockBatch(tx.(*mongo.Transaction), stockId, []*mongo.StockAggregate{dto})
}


func (a *Adapter) StoreStockBatch(tx port.Transactioner, stockId string, aggs []*entity.StockAggregate) error {

	dtos := make([]*mongo.StockAggregate, 0, len(aggs))

	for _, agg := range aggs {
		dtos = append(dtos, &mongo.StockAggregate{
			EventType: agg.EventType,
			Avg:       agg.Average,
			Min:       agg.Min,
			Max:       agg.Max,
			Start:     agg.Start,
			End:       agg.End,
			StartTime: agg.StartTime,
			EndTime:   agg.EndTime,
		})
	}

	return a.mongo.InsertStockBatch(tx.(*mongo.Transaction), stockId, dtos)
}


func (a *Adapter) CreateStoringStartedLog(ctx context.Context, stockId string) error {

	return a.rdbms.CreateAccessInfo(ctx, rdbms.CreateAccessInfoParams{
		ProductID: stockId,
		Status: "started",
	})
}


func (a *Adapter) CreateStoringStoppedLog(ctx context.Context, stockId string) error {

	return a.rdbms.CreateAccessInfo(ctx, rdbms.CreateAccessInfoParams{
		ProductID: stockId,
		Status: "stopped",
	})
}


func (a *Adapter) CreateStoringFailedLog(ctx context.Context, stockId string) error {

	return a.rdbms.CreateAccessInfo(ctx, rdbms.CreateAccessInfoParams{
		ProductID: stockId,
		Status: "failed",
	})
}
