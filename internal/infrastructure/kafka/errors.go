package kafka

import "errors"

var ErrTimeoutRequired = errors.New("timeout setting on ctx required")

var ErrFatalWhileDeletingTopic = errors.New("fatal error while deleting topic")

var ErrTrivalWhileDeletingTopic = errors.New("trival error while deleting topic")

var ErrFailedToDeliveryData = errors.New("failed to delivery data")

var ErrFailedToDeserializeMessage = errors.New("failed to deserialize message")

var ErrFailedToReadData = errors.New("data not arrived")
