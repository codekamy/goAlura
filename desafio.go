// CRIAR UM PROGRAMA QUE FAÇA A CONVERSÃO DE ESCALAS TERMOMÉTRICAS

// UM PROFESSOR DE ENSINO MÉDIO SOLICITOU AOS SEUS ALUNOS QUE DESENVOLVESSEM UM PROGRAMA PARA CONVERTER O VALOR
// DO PONTO DE EBULIÇÃO DA ÁGUA DE KELVIN PARA GRAUS CELSIUS.

// C = K - 273

package main

import "fmt"

const ebulicaoK = 373.0

func main() {
	var ebulicaoC = ebulicaoK - 273.0
	fmt.Println("O valor de ebulição da água em celsius é :", ebulicaoC, "ºC")
}
