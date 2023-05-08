package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNReinasUnicas01(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(1)

	assert.Equal(t, 1, nReinasUnicas.Resolver())
}

func TestNReinasUnicas02(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(2)

	assert.Equal(t, 0, nReinasUnicas.Resolver())
}

func TestNReinasUnicas03(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(3)

	assert.Equal(t, 0, nReinasUnicas.Resolver())
}

func TestNReinasUnicas04(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(4)

	assert.Equal(t, 1, nReinasUnicas.Resolver())
}

func TestNReinasUnicas05(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(5)

	assert.Equal(t, 2, nReinasUnicas.Resolver())
}

func TestNReinasUnicas06(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(6)

	assert.Equal(t, 1, nReinasUnicas.Resolver())
}

func TestNReinasUnicas07(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(7)

	assert.Equal(t, 6, nReinasUnicas.Resolver())
}

func TestNReinasUnicas08(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(8)

	assert.Equal(t, 12, nReinasUnicas.Resolver())
}

func TestNReinasUnicas09(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(9)

	assert.Equal(t, 46, nReinasUnicas.Resolver())
}

func TestNReinasUnicas10(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(10)

	assert.Equal(t, 92, nReinasUnicas.Resolver())
}

func TestNReinasUnicas11(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(11)

	assert.Equal(t, 341, nReinasUnicas.Resolver())
}

func TestNReinasUnicas12(t *testing.T) {
	nReinasUnicas := NewNReinasUnicas(12)

	assert.Equal(t, 1787, nReinasUnicas.Resolver())
}
