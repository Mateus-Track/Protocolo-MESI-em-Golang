package componentes

import (
	"fmt"
	"time"
)

type MP struct { //pelo menos 50 posições;
	Livros [50]Livro
}

func PreencherLivros() MP {
	secoes := []string{
		"Tecnologia",
		"Matemática",
		"História",
		"Literatura",
		"Filosofia",
		"Ciência",
		"Arte",
		"Geografia",
		"Economia",
		"Psicologia",
	}

	var mp MP
	for i := 0; i < 50; i++ {
		secao := secoes[i/5]
		nome := fmt.Sprintf("Livro %d", i+1)
		mp.Livros[i] = Livro{
			Reservas: [][2]time.Time{},
			Nome:     nome,
			Secao:    secao,
		}
	}
	return mp
}
