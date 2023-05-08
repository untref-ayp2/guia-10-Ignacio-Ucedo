package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNReinas01(t *testing.T) {
	nReinas := NewNReinas(1)

	assert.Equal(t, 1, nReinas.Resolver())
}

func TestNReinas02(t *testing.T) {
	nReinas := NewNReinas(2)

	assert.Equal(t, 0, nReinas.Resolver())
}

func TestNReinas03(t *testing.T) {
	nReinas := NewNReinas(3)

	assert.Equal(t, 0, nReinas.Resolver())
}

func TestNReinas04(t *testing.T) {
	nReinas := NewNReinas(4)

	assert.Equal(t, 2, nReinas.Resolver())
}

func TestNReinas05(t *testing.T) {
	nReinas := NewNReinas(5)

	assert.Equal(t, 10, nReinas.Resolver())
}

func TestNReinas06(t *testing.T) {
	nReinas := NewNReinas(6)

	assert.Equal(t, 4, nReinas.Resolver())
}

func TestNReinas07(t *testing.T) {
	nReinas := NewNReinas(7)

	assert.Equal(t, 40, nReinas.Resolver())
}

func TestNReinas08(t *testing.T) {
	nReinas := NewNReinas(8)

	assert.Equal(t, 92, nReinas.Resolver())
}

func TestNReinas09(t *testing.T) {
	nReinas := NewNReinas(9)

	assert.Equal(t, 352, nReinas.Resolver())
}

func TestNReinas10(t *testing.T) {
	nReinas := NewNReinas(10)

	assert.Equal(t, 724, nReinas.Resolver())
}

func TestNReinas11(t *testing.T) {
	nReinas := NewNReinas(11)

	assert.Equal(t, 2680, nReinas.Resolver())
}

func TestNReinas12(t *testing.T) {
	nReinas := NewNReinas(12)

	assert.Equal(t, 14200, nReinas.Resolver())
}
