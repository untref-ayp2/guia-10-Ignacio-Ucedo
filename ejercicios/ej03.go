package ejercicios

type Item struct {
	valor int
	peso  int
}

type Mochila struct {
}

func NewMochila(items []Item, capacidad int) *Mochila {
	return &Mochila{}
}

func (mochila Mochila) Resolver() int {
	return 0
}
