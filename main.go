package goarea

import "math"

//Sempre que a função for pública, é obrigado a ter comentário

// Pi = perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ é responsável por calcular a área de da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi // potência

}

//Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nã visível pois tem "anderlaine" na frente, diferente dos demais que está maiúscula tornando a pública

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
