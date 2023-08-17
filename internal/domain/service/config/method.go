package config

import (
	"context"

	"github.com/Goboolean/fetch-server/internal/domain/vo"
)

func (m *ConfigurationManager) SetStockRelayableTrue(ctx context.Context, stockId string) error {
	return m.relayer.FetchStock(ctx, stockId)
}

func (m *ConfigurationManager) SetStockRelayableFalse(ctx context.Context, stockId string) error {
	return m.relayer.StopFetchingStock(ctx, stockId)
}

func (m *ConfigurationManager) SetStockStoreableTrue(ctx context.Context, stockId string) error {
	return m.persistence.SubscribeRelayer(ctx, stockId)
}

func (m *ConfigurationManager) SetStockStoreableFalse(ctx context.Context, stockId string) error {
	return m.persistence.UnsubscribeRelayer(stockId)
}

func (m *ConfigurationManager) SetStockTransmittableTrue(ctx context.Context, stockId string) error {
	return m.transmitter.SubscribeRelayer(ctx, stockId)
}

func (m *ConfigurationManager) SetStockTransmittableFalse(ctx context.Context, stockId string) error {
	return m.transmitter.UnsubscribeRelayer(stockId)
}


func (m *ConfigurationManager) GetStockConfiguration(ctx context.Context, stockId string) (vo.StockConfiguration, error) {

	tx, err := m.tx.Transaction(context.Background())
	defer tx.Rollback()
	if err != nil {
		return vo.StockConfiguration{}, err
	}

	exists, err := m.db.CheckStockExists(tx, stockId)
	if err != nil {
		return vo.StockConfiguration{}, err
	}

	if !exists {
		return vo.StockConfiguration{}, ErrStockNotFound
	}
	// check stock exist
	// reflect stock info to vo


	if err := tx.Commit(); err != nil {
		return vo.StockConfiguration{}, err
	}

	if isRelayable := m.relayer.IsStockRelayable(stockId); !isRelayable {
		return vo.StockConfiguration{
			StockId: stockId,
			Relayable: false,
			Storeable: false,
			Transmittable: false,
		}, nil
	}

	isStoreable := m.persistence.IsStockStoreable(stockId)
	isTransmittable := m.transmitter.IsStockTransmittable(stockId)

	return vo.StockConfiguration{
		StockId: stockId,
		Relayable: true,
		Storeable: isStoreable,
		Transmittable: isTransmittable,
	}, nil
}


func (m *ConfigurationManager) GetAllStockConfiguration(ctx context.Context) ([]vo.StockConfiguration, error) {

	tx, err := m.tx.Transaction(context.Background())
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	metaList, err := m.db.GetAllStockMetadata(tx)
	if err != nil {
		return nil, err
	}

	confList := make([]vo.StockConfiguration, 0)

	for _, meta := range metaList {
		stockId := meta.StockID
		
		if !m.relayer.IsStockRelayable(stockId) {
			continue
		}

		isStoreable := m.persistence.IsStockStoreable(stockId)
		isTransmittable := m.transmitter.IsStockTransmittable(stockId)

		conf := vo.StockConfiguration{
			StockId: stockId,
			Relayable: true,
			Storeable: isStoreable,
			Transmittable: isTransmittable,
		}

		confList = append(confList, conf)
	}

	return confList, nil
}