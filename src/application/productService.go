package application

import "errors"

type order struct {
	Identifier string
	Products   []string
}

func NewOrder(identifier string, products []string) (*order, error) {
	orderNew := order{Identifier: identifier, Products: products}
	err := orderNew.Verify()
	if err != nil {
		return nil, err
	}
	return &orderNew, nil
}

func (order order) Verify() error {
	if order.Identifier == "" {
		return errors.New("error: identifier empty")
	}

	if order.Products == nil {
		return errors.New("error: products nil")
	}

	if len(order.Products) == 0 {
		return errors.New("error: products empty")
	}
	if len(order.Products) > 50 {
		return errors.New("error: Product list cannot be longer than 50")
	}

	return nil
}

func (order order) TotalProducts() int {
	return len(order.Products)
}
