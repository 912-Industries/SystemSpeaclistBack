package main
// Author Yuri Bueno 0623221
import (
	"SpecialistSystem/controllers"
	"fmt"
)

func main() {
	sistema := controllers.NovoSistemaEspecialista()

	var escolha int
	var genero int

	fmt.Println("Digite 1 - Receita para Homem\nDigite 2 - Receita para Mulher")
	_, err := fmt.Scan(&genero)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Digite 1 - Receita para emagrecer\nDigite 2 - Receita para engordar")
	_, err = fmt.Scan(&escolha)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if escolha == 1 && genero == 1 {
		for i := range sistema.PerdaPesoHomem {
			fmt.Printf("Horario: %s\n", sistema.PerdaPesoHomem[i].Horario)
			fmt.Printf("Descricao: %s\n", sistema.PerdaPesoHomem[i].Descricao)
		}
	}
	if escolha == 2 && genero == 1 {
		for i := range sistema.GanhoPesoHomem {
			fmt.Printf("Horario: %s\n", sistema.GanhoPesoHomem[i].Horario)
			fmt.Printf("Descricao: %s\n", sistema.GanhoPesoHomem[i].Descricao)
		}
		fmt.Println("You entered:", escolha)

	}
	if escolha == 1 && genero == 2 {
		for i := range sistema.PerdaPesoMulher {
			fmt.Printf("Horario: %s\n", sistema.PerdaPesoMulher[i].Horario)
			fmt.Printf("Descricao: %s\n", sistema.PerdaPesoMulher[i].Descricao)
		}
	}
	if escolha == 2 && genero == 2 {
		for i := range sistema.GanhoPesoMulher {
			fmt.Printf("Horario: %s\n", sistema.GanhoPesoMulher[i].Horario)
			fmt.Printf("Descricao: %s\n", sistema.GanhoPesoMulher[i].Descricao)
		}
		fmt.Println("You entered:", escolha)

	}
}
