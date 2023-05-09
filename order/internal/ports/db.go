package ports 

import "github.com/jfelipeforero/microservices/order/internal/application/core/domain"

type DBPort interface {
  Get(id string) (domain.Order, error)
  Save(*domain.Order) error
}
