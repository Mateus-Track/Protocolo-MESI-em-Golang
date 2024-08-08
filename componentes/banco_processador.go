package componentes

type BancoProcessadores struct {
	BP []Processador
}

func InicializaBP(QUANTIDADE_USUARIOS int) BancoProcessadores {
	return BancoProcessadores{
		BP: make([]Processador, QUANTIDADE_USUARIOS),
	}
}
