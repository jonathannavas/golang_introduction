package structs

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (p Product) GetTotalPrice() float64 {
	return float64(p.Count) * (p.Price)
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

type Cart struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (c *Cart) AddProducs(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Cart) Total() float64 {
	var total float64
	for _, c := range c.Products {
		total += c.GetTotalPrice()
	}
	return total
}

func NewCart(userID uint) Cart {
	return Cart{UserID: userID}
}
