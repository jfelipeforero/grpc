package grpc

import (
  "context"
  "github.com/jfelipeforero/microservices-proto/golang/order"
  "github.com/jfelipeforero/grpc/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest)
        (*order.CreateOrderResponse, error) {
    var orderItems []domain.orderItems
    for _, orderItem := range request.OrderItem {
        orderItems = append(orderItems, domain.OrderItem{
            ProductCode: orderItem.ProductCode,
            UnitPrice: orderItem.UnitPrice,
            Quantity: orderItem.Quantity,
        }) 
    }
    newOrder := domain.NewOrder(request.UserId, orderItems)
    result, err := a.api.PlaceOrder(newOrder)
    if err != nil {
      return nil, err
    }
    return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
