package buycycle

import (
	"log"

	"github.com/Goboolean/fetch-server/internal/infrastructure/ws"
)



func (s *Subscriber) SubscribeStockAggs(stocks ...string) error {
	for _, stock := range stocks {
		if err := s.WriteJSON(stock); err != nil {
			return err
		}
	}

	return nil
}


func RelayMessageToReceiver(s *Subscriber) {
	var data StockAggregate

	for {
		select {
		case <- s.ctx.Done():
			return
		default:

			if err := s.ReadJSON(&data); err != nil {
				log.Fatalf("failed to read json data: %v", err)
				continue
			}

			stockAggs := &ws.StockAggregate{}

			if err := s.r.OnReceiveStockAggs(stockAggs); err != nil {
				log.Fatalf("failed to process data: %v", err)
				continue
			}
		}
	}
}



func (s *Subscriber) UnsubscribeStockAggs(stock ...string) error {
	return nil
}