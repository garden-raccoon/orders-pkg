package models

import (
	proto "github.com/garden-raccoon/orders-pkg/protocols/orders-pkg"
	"github.com/gofrs/uuid"
	"time"
)

type Order struct {
	OrderUuid   uuid.UUID `json:"order_uuid"`
	Params      []*Param  `json:"params"`
	UserUuid    uuid.UUID `json:"user_uuid"`
	OrderStatus int       `json:"order_status"`
}
type Param struct {
	MealUuid     uuid.UUID `json:"meal_uuid"`
	DeliveryUuid uuid.UUID `json:"delivery_uuid"`
	Quantity     int       `json:"quantity"`
	Date         time.Time `json:"date"`
}

// Proto is
func Proto(o *Order) *proto.Order {
	return &proto.Order{
		OrderUuid:   o.OrderUuid.Bytes(),
		UserUuid:    o.UserUuid.Bytes(),
		OrderStatus: int64(o.OrderStatus),
		Params:      ParamsToProto(o.Params),
	}
}

func OrderFromProto(pb *proto.Order) *Order {
	return &Order{
		OrderUuid:   uuid.FromBytesOrNil(pb.OrderUuid),
		Params:      ParamsFromProto(pb.Params),
		UserUuid:    uuid.FromBytesOrNil(pb.UserUuid),
		OrderStatus: int(pb.OrderStatus),
	}
}

func OrdersToProto(orders []*Order) (*proto.Orders, error) {
	pb := &proto.Orders{}

	for _, b := range orders {
		pb.Orders = append(pb.Orders, Proto(b))
	}
	return pb, nil
}

func OrdersFromProto(pb *proto.Orders) ([]*Order, error) {
	var orders []*Order
	for _, b := range pb.Orders {

		orders = append(orders, OrderFromProto(b))
	}
	return orders, nil
}
