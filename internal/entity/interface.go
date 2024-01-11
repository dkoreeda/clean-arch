package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int32, error)
	List(limit, offset int32) ([]Order, error)
}
