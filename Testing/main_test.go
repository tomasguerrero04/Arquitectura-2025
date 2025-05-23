package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	resultado := suma(8, 4)

	assert.Equal(t, resultado, 12, "La division no es correcta")

}
