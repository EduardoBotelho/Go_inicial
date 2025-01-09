// declarar um pacote principal
package main

//importar funcao fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoF = 212.0

func main() {
	 tempF := ebulicaoF            // variavel do valor da temperatura em graus F
	 tempC := (tempF - 32) * 5 / 9 // Conversao de F para C
	//aparece o resultado na tela
	fmt.Println("A temperatura de ebulicao da agua em F = ", tempF)
	fmt.Println("A temperatura de ebulicao da agua em C = ", tempC)
}
