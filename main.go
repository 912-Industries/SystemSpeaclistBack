package main

type Receita struct {
	Horario      string
	Nome         string
	Descricao    string
	Calorias     int
	Proteinas    int
	Carboidratos int
	Gorduras     int
}

type SistemaEspecialista struct {
	PerdaPeso []Receita
	GanhoPeso []Receita
}

func NovoSistemaEspecialista() *SistemaEspecialista {
	return &SistemaEspecialista{
		GanhoPeso: []Receita{
			{
				Nome:         "4 Ovos cozidos",
				Descricao:    "Ovos cozidos para ganho de peso.",
				Calorias:     256,
				Proteinas:    24,
				Carboidratos: 4,
				Gorduras:     16,
			},
			{
				Nome:         "2 Big Mac",
				Descricao:    "2 Big Mac de leve.",
				Calorias:     1000,
				Proteinas:    52,
				Carboidratos: 84,
				Gorduras:     50,
			},
			{
				Nome:         "2 Whopper Rodeio",
				Descricao:    "Whopper pra ganho de peso",
				Calorias:     1530,
				Proteinas:    40,
				Carboidratos: 169,
				Gorduras:     78,
			},
		},
		PerdaPeso: []Receita{
			{
				Nome:         "2 Folhas de alface",
				Calorias:     2,
				Proteinas:    0,
				Carboidratos: 1,
				Gorduras:     0,
			},
			{
				Nome:         "4 Ovos cozidos",
				Descricao:    "Ovos cozidos para perda de peso.",
				Calorias:     256,
				Proteinas:    24,
				Carboidratos: 4,
				Gorduras:     16,
			},
			{
				Nome:         "2 Folhas de alface",
				Calorias:     2,
				Proteinas:    0,
				Carboidratos: 1,
				Gorduras:     0,
			},
		},
	}
}

func main() {

}
