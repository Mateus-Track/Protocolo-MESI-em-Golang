package componentes

type Processador struct {
	Cachezinha Cache //precisa ser maiúsculo o nome da variável..
}

func InicializaProcessador() Processador {
	return Processador{
		Cachezinha: InicializaCache(),
	}
}
