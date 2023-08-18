package app

import "github.com/akramulfata10/gotoko/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() [] Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Order{}},
		{Model: models.Category{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderCustomer{}},
		{Model: models.Payment{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Shipment{}},
	}
}
