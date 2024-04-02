package main

import (
	"SpecialistSystem/controllers"
	"fmt"
)
func main() {
	sistema := controllers.NovoSistemaEspecialista()

	for i := range sistema.PerdaPeso {
		fmt.Printf("Nome: %s\n", sistema.PerdaPeso[i].Nome)
		fmt.Printf("Descricao: %s\n", sistema.PerdaPeso[i].Descricao)
		fmt.Printf("Calorias: %d\n", sistema.PerdaPeso[i].Calorias)
		fmt.Printf("Proteinas: %d\n", sistema.PerdaPeso[i].Proteinas)
		fmt.Printf("Carboidratos: %d\n", sistema.PerdaPeso[i].Carboidratos)
		fmt.Printf("Gorduras: %d\n", sistema.PerdaPeso[i].Gorduras)
	}
}
