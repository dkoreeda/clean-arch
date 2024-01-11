package usecase

import (
	"fmt"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderListInputDTO struct {
	Limit  int32
	Offset int32
}

type OrderListOutputDTO struct {
	Orders []entity.Order
	Total  int32
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (l *ListOrdersUseCase) Execute() (OrderListOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		fmt.Printf("execute error: %v\n", err)
		return OrderListOutputDTO{}, err
	}

	total, err := l.OrderRepository.GetTotal()
	if err != nil {
		return OrderListOutputDTO{}, err
	}

	return OrderListOutputDTO{
		Orders: orders,
		Total:  total,
	}, nil
}
