package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da adivinhação")
	fmt.Println("Um número aleatória será sorteado. Tente acertar.  O número inteiro entre 0 e 100")
	x := rand.Int64N(101)
	scanner := bufio.NewScanner((os.Stdin))
	chutes := [10]int64{}
	for i := range chutes {
		fmt.Print("Qual seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro")
			return
		}
		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que ", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que ", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns, você acertou! O numero era %d\n"+
					"Voce precisou de %d tentativas.\n"+
					"Essas foram as suas tentativas: %v\n ", x, i+1, chutes[:i+1])
			return
		}
		chutes[i] = chuteInt
	}
	fmt.Printf("Infelizmente voce nao acertou o numero, que era: %d\n"+
		"Voce teve 10 tentativas.\n"+
		"Essas foram as suas tentativas: %v\n ", x, chutes)
}
