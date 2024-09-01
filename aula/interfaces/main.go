package main

import "fmt"

//Interface para representar veiculos
//VelocidadeMax, meioLocomocao

type Veiculo interface {
	VelocidadeMax() int
	MeioLocomocao() string
	MostrarNome() string
}

type Carro struct {
	Nome             string
	MaximaVelocidade int
}

type Trem struct {
	Nome             string
	MaximaVelocidade int
}

func (c Carro) VelocidadeMax() int {
	return c.MaximaVelocidade
}

func (c Carro) MeioLocomocao() string {
	return "Rodas"
}

func (c Carro) MostrarNome() string {
	return c.Nome
}

func (t Trem) VelocidadeMax() int {
	return t.MaximaVelocidade
}

func (t Trem) MeioLocomocao() string {
	return "Trilhos"
}

func (t Trem) MostrarNome() string {
	return t.Nome
}

func DescreverVeiculo(v Veiculo) {
	fmt.Println("Veiculo Atual:", v.MostrarNome())
	fmt.Printf("Meio de locomoção: %s, Velocidade Máxima: %d km/h \n", v.MeioLocomocao(), v.VelocidadeMax())
}

func main() {
	meuCarro := Carro{Nome: "Gol 1.0", MaximaVelocidade: 130}
	meuTrem := Trem{Nome: "Trem GO", MaximaVelocidade: 150}

	DescreverVeiculo(meuCarro)
	DescreverVeiculo(meuTrem)
}
