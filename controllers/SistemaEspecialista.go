package controllers

import "SpecialistSystem/models"

var receita models.Receita

type SistemaEspecialista struct {
	PerdaPeso []models.Receita
	GanhoPeso []models.Receita
}

func NovoSistemaEspecialista() *SistemaEspecialista {
	return &SistemaEspecialista{
		GanhoPeso: []models.Receita{
			{
				Horario:      "8:00",
				Nome:         "4 Ovos cozidos",
				Descricao:    "Ovos cozidos para ganho de peso.",
				Calorias:     256,
				Proteinas:    24,
				Carboidratos: 4,
				Gorduras:     16,
			},
			{
				Horario:      "12:00",
				Nome:         "2 Big Mac",
				Descricao:    "2 Big Mac de leve.",
				Calorias:     1000,
				Proteinas:    52,
				Carboidratos: 84,
				Gorduras:     50,
			},
			{
				Horario:      "20:00",
				Nome:         "2 Whopper Rodeio",
				Descricao:    "Whopper pra ganho de peso",
				Calorias:     1530,
				Proteinas:    40,
				Carboidratos: 169,
				Gorduras:     78,
			},
		},
		PerdaPeso: []models.Receita{
			{
				Horario:      "8:00",
				Nome:         "2 Folhas de alface",
				Descricao:    "2 folhas de alface para emagrrecer",
				Calorias:     2,
				Proteinas:    0,
				Carboidratos: 1,
				Gorduras:     0,
			},
			{
				Horario:      "12:00",
				Nome:         "4 Ovos cozidos",
				Descricao:    "Ovos cozidos para perda de peso.",
				Calorias:     256,
				Proteinas:    24,
				Carboidratos: 4,
				Gorduras:     16,
			},
			{
				Horario:      "20:00",
				Nome:         "2 Folhas de alface",
				Descricao:    "2 folhas de alface para emagrrecer",
				Calorias:     2,
				Proteinas:    0,
				Carboidratos: 1,
				Gorduras:     0,
			},
		},
	}
}
