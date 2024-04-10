package controllers

import "SpecialistSystem/models"

var receita models.Receita

type SistemaEspecialista struct {
	PerdaPesoHomem  []models.Receita
	GanhoPesoHomem  []models.Receita
	PerdaPesoMulher []models.Receita
	GanhoPesoMulher []models.Receita
}

func NovoSistemaEspecialista() *SistemaEspecialista {
	return &SistemaEspecialista{
		PerdaPesoHomem: []models.Receita{
			{
				Horario:   "8:00",
				Descricao: "Omelete de claras de ovos com vegetais (tomate, espinafre, cogumelos) e uma fatia de pão integral. 1 xícara de frutas frescas (morangos, melão, kiwi).",
			},
			{
				Horario:   "12:00",
				Descricao: "Salada de quinoa com legumes grelhados (pimentão, abobrinha, berinjela) e peito de frango grelhado. 1 porção de iogurte grego natural com 1/4 xícara de nozes.",
			},
			{
				Horario:   "15:00",
				Descricao: "Filé de salmão grelhado com aspargos e quinoa. Salada verde com vinagrete balsâmico",
			},
			{
				Horario:   "20:00",
				Descricao: "1 maçã ou pera com 1 colher de sopa de manteiga de amendoim natural.",
			},
			{
				Horario:   "22:00",
				Descricao: "1 copo de leite desnatado ou leite de amêndoa.",
			},
		},
		GanhoPesoHomem: []models.Receita{
			{
				Horario:   "8:00",
				Descricao: "Omelete com três ovos inteiros, queijo e vegetais (como espinafre, cogumelos e tomate) Duas fatias de pão integral torrado com manteiga de amendoim Um copo de suco de frutas natural",
			},
			{
				Horario:   "12:00",
				Descricao: "200g de peito de frango grelhado 1 xícara de quinoa cozida Vegetais cozidos no vapor (brócolis, cenoura, couve-flor) Salada verde com abacate, nozes e azeite de oliva",
			},
			{
				Horario:   "15:00",
				Descricao: "1 copo de shake de proteína com leite integral, banana, aveia e uma colher de sopa de manteiga de amendoim 1 maçã",
			},
			{
				Horario:   "20:00",
				Descricao: "200g de salmão assado 1 xícara de arroz integral Aspargos grelhados com azeite de oliva e alho Salada de rúcula com tomate, queijo feta e nozes",
			},
			{
				Horario:   "22:00",
				Descricao: "1 copo de iogurte grego natural com uma colher de sopa de mel e algumas frutas secas (como damascos ou figos).",
			},
		},

		PerdaPesoMulher: []models.Receita{
			{
				Horario:   "8:00",
				Descricao: "Smoothie verde com espinafre, abacate, banana e leite de amêndoa. 1 colher de sopa de semente de chia ou linhaça.",
			},
			{
				Horario:   "12:00",
				Descricao: "Salada de quinoa com legumes grelhados (pimentão, abobrinha, berinjela) e peito de frango grelhado. 1 porção de iogurte grego natural com 1/4 xícara de nozes.",
			},
			{
				Horario:   "15:00",
				Descricao: "Filé de salmão grelhado com aspargos e quinoa. Salada verde com vinagrete balsâmico",
			},
			{
				Horario:   "20:00",
				Descricao: ".Frango ao curry com legumes (brócolis, cenoura, couve-flor) e arroz integral.",
			},
			{
				Horario:   "22:00",
				Descricao: "1 copo de leite desnatado ou leite de amêndoa.",
			},
		},
		GanhoPesoMulher: []models.Receita{
			{
				Horario:   "8:00",
				Descricao: "Smoothie com 1 banana, 1/2 xícara de morangos, 1/2 xícara de aveia em flocos, 1 colher de sopa de manteiga de amendoim e leite de amêndoa 2 fatias de pão integral com abacate amassado e um ovo pochê por cima",
			},
			{
				Horario:   "12:00",
				Descricao: "150g de tofu grelhado com tempero de ervas 1/2 xícara de quinoa cozida Salada de folhas verdes com cenoura ralada, sementes de abóbora e vinagrete de limão.",
			},
			{
				Horario:   "15:00",
				Descricao: "1 pequeno punhado de castanhas e nozes mistas 1 pera",
			},
			{
				Horario:   "20:00",
        			Descricao: "150g de peito de frango grelhado 1 batata-doce assada com um fio de azeite de oliva e alecrim Brócolis cozidos no vapor com limão e pimenta",
			},
			{
				Horario:   "22:00",
        			Descricao: "1 xícara de leite integral com 1 colher de sopa de cacau em pó e 1 colher de chá de mel",
			},
		},
	}
}
