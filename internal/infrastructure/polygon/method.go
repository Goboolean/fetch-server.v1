package polygon

import (
	"log"

	polygonws "github.com/polygon-io/client-go/websocket"
	"github.com/polygon-io/client-go/websocket/models"
)





func (s *Subscriber) SubscribeStocksSecAggs(stock string) error {

	if err := s.conn.Subscribe(polygonws.StocksSecAggs, stock); err != nil {
		return err
	}

	return nil
}


func RelayMessageToReceiver(c *Subscriber) {
	for {
		select {
		case <- c.ctx.Done():
			return
		case err := <-c.conn.Error():
			log.Fatal("error: ", err)
			return
		case out, more := <-c.conn.Output():
			if !more {
				return
			}

			if err := c.r.OnReceivePolygonStockAggs(out.(models.EquityAgg)); err != nil {
				log.Fatal(err)
			}
		}
	}
}