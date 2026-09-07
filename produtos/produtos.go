package produtos

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Code     int     `json:"code"`
	Quantity int     `json:"quantity"`
}

func (p *Product) TotalValue() float64 {
	return p.Price * float64(p.Quantity)
}