package p1

type Fruit struct {
	Name  string  `json:"name" query:"name"`
	Prict float64 `json:"price" query:"price"`
}

func (Fruit) Sale() (price float64) {
	return 100
}

func (f *Fruit) Sale2() (price float64) {
	return f.Prict * 100
}
