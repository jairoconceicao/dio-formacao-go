package main

import "testing"

func TestShouldSumCorrectly(t *testing.T) {
	result := soma(1, 2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Valor esperado %d, mas obteve %d", expected, result)
	}
}

func TestShouldMultiplyCorrectly(t *testing.T) {
	result := multiplicacao(2, 3, 4)
	expected := 24
	if result != expected {
		t.Errorf("Valor esperado %d, mas obteve %d", expected, result)
	}
}

func TestShouldDivideCorrectly(t *testing.T) {
	result, err := divisao(10, 2)
	expected := 5
	if err != nil {
		t.Errorf("Erro inesperado: %s", err)
	}
	if result != expected {
		t.Errorf("Valor esperado %d, mas obteve %d", expected, result)
	}
}

func TestShouldHandleDivisionByZero(t *testing.T) {
	_, err := divisao(10, 0)
	if err == nil {
		t.Error("Esperava um erro de divisão por zero, mas não recebeu nenhum")
	}
}

func TestShouldSubtractCorrectly(t *testing.T) {
	result := subtracao(10, 3)
	expected := 7
	if result != expected {
		t.Errorf("Valor esperado %d, mas obteve %d", expected, result)
	}
}
