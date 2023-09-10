package models

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Category{}},
		{Model: OrderCustomer{}},
		{Model: OrderItem{}},
		{Model: Order{}},
		{Model: Payment{}},
		{Model: ProductImage{}},
		{Model: Product{}},
		{Model: Section{}},
		{Model: Shipment{}},
	}
}
