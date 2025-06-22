package graphql

import "github.com/lucianocasa/api/internal/order"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service order.Service
}
