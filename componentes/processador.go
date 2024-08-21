package componentes

type Processador struct {
	id         int
	Cachezinha Cache //precisa ser maiúsculo o nome da variável..
}

func InicializaProcessador(id int) Processador {
	return Processador{
		id:         id,
		Cachezinha: InicializaCache(id),
	}
}
