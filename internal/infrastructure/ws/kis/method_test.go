package kis_test

import (
	"os"
	"testing"
	"time"

	"github.com/Goboolean/fetch-server.v1/internal/infrastructure/ws/kis"
)

func Test_Method(t *testing.T) {

	var instance *kis.Subscriber = new(kis.Subscriber)

	t.Run("GetApprovalKey (case:fail)", func(t *testing.T) {

		_, err := instance.GetApprovalKey("", "")
		if err == nil {
			t.Errorf("GetApprovalKey() = %v, want error", err)
			return
		}
	})

	t.Run("GetApprovalKey (case:success)", func(t *testing.T) {

		key, err := instance.GetApprovalKey(os.Getenv("KIS_APPKEY"), os.Getenv("KIS_SECRET"))
		if err != nil {
			t.Errorf("GetApprovalKey() = %v", err)
			return
		}

		if key == "" {
			t.Errorf("GetApprovalKey() = %v, want not empty", key)
			return
		}
	})
}

func Test_SubscribeStockAggs(t *testing.T) {

	const (
		symbol      = "DNASAAPL"
		falseSymbol = "FALSE"
	)

	var (
		countBeforeSubscription  int
		countAfterSubscription   int
		countAfterUnsubscription int
	)

	t.Run("FalseSubscribe", func(t *testing.T) {
		if err := instance.SubscribeStockAggs(falseSymbol); err == nil {
			t.Errorf("SubscrbeStockAggs() = %v, want error", err)
			return
		}
	})

	t.Run("Subscribe", func(t *testing.T) {

		if err := instance.SubscribeStockAggs(symbol); err != nil {
			t.Errorf("SubscrbeStockAggs() = %v", err)
			return
		}

		countBeforeSubscription = count

		time.Sleep(time.Second * 3 / 2)

		countAfterSubscription = count
		diff := countAfterSubscription - countBeforeSubscription

		if diff == 0 {
			t.Errorf("SubscribeStockAggs() received %d, want many", diff)
			return
		}
	})

	t.Run("SubscribeTwice", func(t *testing.T) {
		if err := instance.SubscribeStockAggs(symbol); err == nil {
			t.Errorf("SubscrbeStockAggs() = %v, want error", err)
			return
		}
	})

	t.Run("Unsubscribe", func(t *testing.T) {
		if err := instance.UnsubscribeStockAggs(symbol); err != nil {
			t.Errorf("UnsubscribeStockAggs() = %v", err)
			return
		}

		time.Sleep(1 * time.Second)

		countAfterUnsubscription = count
		diff := countAfterUnsubscription - countAfterSubscription

		if diff != 0 {
			t.Errorf("UnsubscribeStockAggs() received %d, want 0", diff)
			return
		}
	})

	t.Run("UnsubscribeTwice", func(t *testing.T) {
		if err := instance.UnsubscribeStockAggs(symbol); err == nil {
			t.Errorf("UnsubscribeStockAggs() = %v, want error", err)
			return
		}
	})
}
