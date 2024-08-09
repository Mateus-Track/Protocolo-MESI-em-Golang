package componentes

type BancoProcessadores struct {
	BP []Processador
}

func InicializaBP(QUANTIDADE_USUARIOS int) BancoProcessadores {
	bp := BancoProcessadores{
		BP: make([]Processador, QUANTIDADE_USUARIOS),
	}

	// Inicializar cada Processador dentro do Banco
	for i := range bp.BP {
		bp.BP[i] = InicializaProcessador()
	}

	return bp
}
