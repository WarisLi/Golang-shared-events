package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(LowProductQuantityNotificationEvent{}).Name(),
	reflect.TypeOf(NewOrderNotificationEvent{}).Name(),
}

type Event interface {
}

type LowProductQuantityNotificationEvent struct {
	Name     string
	Quantity int
}

type NewOrderNotificationEvent struct {
	OrderId         uint
	ProductName     string
	Quantity        int
	CustomerName    string
	CustomerAddress string
}
