package main

type Customer struct {
	ID   string
	Name string
}

type Purchase struct {
	ID               string
	Value            float64
	Customer         []*Customer
	DateTimePurchase string
}

func SetPurchase() *Purchase {
	return &Purchase{
		ID:    "034985kosdjf",
		Value: 65456.54,
		Customer: []*Customer{
			{
				ID:   "456456",
				Name: "Ze"},
		},
		DateTimePurchase: "2023-12-21",
	}
}
