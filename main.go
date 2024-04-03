package main

import (
	"SpecialistSystem/controllers"
	"fmt"
)

func main() {
	sistema := controllers.NovoSistemaEspecialista()

	var escolha int

	fmt.Println("Digite 1 - Receita para emagrecer\nDigite 2 - Receita para engordar")
	_, err := fmt.Scan(&escolha)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if escolha == 1 {
		for i := range sistema.PerdaPeso {
			fmt.Printf("Horario: %s\n", sistema.PerdaPeso[i].Horario)
			fmt.Printf("Nome: %s\n", sistema.PerdaPeso[i].Nome)
			fmt.Printf("Descricao: %s\n", sistema.PerdaPeso[i].Descricao)
			fmt.Printf("Calorias: %d\n", sistema.PerdaPeso[i].Calorias)
			fmt.Printf("Proteinas: %d\n", sistema.PerdaPeso[i].Proteinas)
			fmt.Printf("Carboidratos: %d\n", sistema.PerdaPeso[i].Carboidratos)
			fmt.Printf("Gorduras: %d\n", sistema.PerdaPeso[i].Gorduras)
		}

		if escolha == 2 {
			for i := range sistema.GanhoPeso {
				fmt.Printf("Horario: %s\n", sistema.GanhoPeso[i].Horario)
				fmt.Printf("Nome: %s\n", sistema.GanhoPeso[i].Nome)
				fmt.Printf("Descricao: %s\n", sistema.GanhoPeso[i].Descricao)
				fmt.Printf("Calorias: %d\n", sistema.GanhoPeso[i].Calorias)
				fmt.Printf("Proteinas: %d\n", sistema.GanhoPeso[i].Proteinas)
				fmt.Printf("Carboidratos: %d\n", sistema.GanhoPeso[i].Carboidratos)
				fmt.Printf("Gorduras: %d\n", sistema.GanhoPeso[i].Gorduras)
			}
			fmt.Println("You entered:", escolha)

		}
	}
}
