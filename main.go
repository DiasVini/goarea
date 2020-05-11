package goarea

import "math"

const Pi = 3.1416

// Circ é responsavel por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcuar a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Essa não é uma função visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
