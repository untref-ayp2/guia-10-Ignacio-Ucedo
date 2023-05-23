package ejercicios

import "fmt"

//solo funciona 1 2 y 3 en test
type NReinas struct {
	tablero [][]bool
}

func NewNReinas(tam int) *NReinas {

	tablero := make([][]bool, tam)

	for i := range tablero {
		tablero[i] = make([]bool, tam)
	}

	return &NReinas{tablero: tablero}
}

func (nReinas NReinas) Resolver() int {
	var soluciones int
	var pasos []int // estos números almacenados significan las columnas en las que coloco cada reina.
	//yendo desde la primera fila

	for i := 0; i < len(nReinas.tablero); i++ {
		for j := 0; j < len(nReinas.tablero[0]); j++ {

			if nReinas.tablero[i][j] { //si está ocupado (porque volví en mis pasos) la desocupo
				//para pasar a la sig. columna
				nReinas.tablero[i][j] = false
				fmt.Printf("Saqué la de (%v , %v)\n", i, j)
				pasos = pasos[:len(pasos)-1]
				if j == len(nReinas.tablero[0])-1 { // si la que saqué ya está en la última columna
					//ubico la anterior en otra columna (siempre y cuando no esté parado en la primera fila)
					if i > 0 {
						i -= 2
					} else { // si estoy en la primera fila se terminaron las posibles soluciones
						return soluciones
					}

				}
			} else if esFactible(i, j, pasos) { // si es factible colocarla en la columna j
				nReinas.tablero[i][j] = true // La coloco (aclaración: falso: libre , true: ocupado)
				fmt.Printf("Puse una en (%v , %v)\n", i, j)
				pasos = append(pasos, j)
				if i == len(nReinas.tablero)-1 { // si coloqué la última
					soluciones++
					fmt.Print("\nSumé una solución\n\n")
					if i > 0 {
						i -= 2 //vuelvo en mis pasos siempre y cuando no esté en la primera fila (fila 0)
					}
				}

				break
			}

			if j == len(nReinas.tablero[0])-1 { // si no la pude ubicar en ninguna columna
				//ubico la anterior en otra columna
				if i > 0 {
					i -= 2
				}

			}
		}
	}

	return soluciones
}

func esFactible(fila int, columna int, reinas []int) bool {

	for i := 0; i < len(reinas); i++ {
		if reinas[i] == columna || fila+columna == i+reinas[i] || fila-columna == i-reinas[i] {
			return false
		}
	}

	return fila != len(reinas)-1 //si estoy parado en la misma fila de una que ya coloqué
	// return true
}
