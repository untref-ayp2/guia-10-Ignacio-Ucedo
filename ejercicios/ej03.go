package ejercicios

import (
	"math"
)

type Item struct {
	valor int
	peso  int
}

type Mochila struct {
	tablero [][]int
	items   []Item
}

func NewMochila(items []Item, capacidad int) *Mochila {

	tablero := make([][]int, len(items))

	for i := range tablero {
		tablero[i] = make([]int, capacidad)
	}

	return &Mochila{tablero: tablero, items: items}
}

func (mochila Mochila) Resolver() int {

	for i := 0; i < len(mochila.tablero); i++ {
		for j := 0; j < len(mochila.tablero[i]); j++ {
			capacidadActual := j + 1
			if i == 0 {
				if mochila.items[0].peso <= capacidadActual {
					mochila.tablero[i][j] = mochila.items[0].valor
				}
			} else {

				noIncluyoItem := mochila.tablero[i-1][j]
				var incluyoItem int

				if mochila.items[i].peso <= capacidadActual {
					capacidadRestante := capacidadActual - mochila.items[i].peso
					if capacidadRestante == 0 {
						incluyoItem = mochila.items[i].valor
					} else {
						incluyoItem = mochila.items[i].valor + mochila.tablero[i-1][capacidadRestante-1]
					}

				}
				mochila.tablero[i][j] = int(math.Max(float64(incluyoItem), float64(noIncluyoItem)))
			}
		}
	}

	return mochila.tablero[len(mochila.items)-1][len(mochila.tablero[0])-1]
}

//https://www.youtube.com/watch?v=8LusJS5-AGo
//la columna 0 es innecesaria
