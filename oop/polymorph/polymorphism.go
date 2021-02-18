package polymorphism

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

func NewItem(name string, price float64, quantity int) *Item {
	if price <= 0 || len(name) == 0 {
		return nil
	}
	return &Item{name, price, quantity}
}

func (t *Item) GetName() string {
	return t.Name
}

func (t *Item) SetName(n string) {
	if len(n) != 0 {
		t.Name = n
	}
}

func (t *Item) GetPrice() float64 {
	return t.Price
}

func (t *Item) SetPrice(p float64) {
	if p > 0 {
		t.Price = p
	}
}

func (t *Item) GetQuantity() int {
	return t.Quantity
}

func (t *Item) SetQuantity(q int) {
	if q > 0 {
		t.Quantity = q
	}
}

func (t *Item) Cost() float64 {
	return t.Price * float64(t.Quantity)
}
