package testunitario

import "testing"

func TestSumaSimple(t *testing.T) {
	total := Suma(5, 5)
	esperado := 10
	if total != esperado {
		t.Errorf("Suma incorrecta, obtuvo %d y se esperaba %d", esperado, 10)
	}
}

func TestSumaTabla(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, elementos := range tabla {
		total := Suma(elementos.a, elementos.b)
		if total != elementos.c {
			t.Errorf("Suma incorrecta, obtuvo %d y se esperaba %d", total, elementos.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{6, 4, 6},
		{25, 30, 30},
	}

	for _, elementos := range tabla {
		max := GetMax(elementos.a, elementos.b)
		if max != elementos.c {
			t.Errorf("Máximo incorrecto, obtuvo %d y se esperaba %d", max, elementos.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		a int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, elementos := range tabla {
		fibonacci := Fibonacci(elementos.a)
		if fibonacci != elementos.r {
			t.Errorf("Fibonacci incorrecto, obtuvo %d y se esperaba %d", fibonacci, elementos.r)
		}
	}
}
