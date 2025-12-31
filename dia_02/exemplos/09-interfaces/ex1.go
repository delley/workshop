package main

import "fmt"

// Geo define o comportamento básico de figuras geométricas.
type Geo interface {
	Area() float64
}

// Retangulo representa um retângulo.
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Area calcula a área de um retângulo.
func (r *Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Triangulo representa um triângulo.
type Triangulo struct {
	Base   float64
	Altura float64
}

// Area calcula a área de um triângulo.
func (t *Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

// imprimeArea recebe qualquer tipo que satisfaça a interface Geo.
func imprimeArea(g Geo) {
	fmt.Printf("Área: %.2f\n", g.Area())
}

func main() {
	r := Retangulo{
		Largura: 5,
		Altura:  10,
	}

	t := Triangulo{
		Base:   10,
		Altura: 5,
	}

	imprimeArea(&r)
	imprimeArea(&t)
}
