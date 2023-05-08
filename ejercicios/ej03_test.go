package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMochilaUno(t *testing.T) {
	items := []Item{
		{60, 10},
		{100, 20},
		{120, 30},
	}
	capacidad := 40
	mochila := NewMochila(items, capacidad)

	assert.Equal(t, 180, mochila.Resolver())
}

func TestMochilaDos(t *testing.T) {
	items := []Item{
		{55, 95},
		{10, 4},
		{47, 60},
		{5, 32},
		{4, 23},
		{50, 72},
		{8, 80},
		{61, 62},
		{85, 65},
		{87, 46},
	}
	capacidad := 269
	mochila := NewMochila(items, capacidad)

	assert.Equal(t, 295, mochila.Resolver())
}

func TestMochilaTres(t *testing.T) {
	items := []Item{
		{44, 92},
		{46, 4},
		{90, 43},
		{72, 83},
		{91, 84},
		{40, 68},
		{75, 92},
		{35, 82},
		{8, 6},
		{54, 44},
		{78, 32},
		{40, 18},
		{77, 56},
		{15, 83},
		{61, 25},
		{17, 96},
		{75, 70},
		{29, 48},
		{75, 14},
		{63, 58},
	}
	capacidad := 878
	mochila := NewMochila(items, capacidad)

	assert.Equal(t, 1024, mochila.Resolver())
}
