package componentes

type Cache struct { //pelo menos 5 posições.
	Linhas [5]Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
	Fila   []uint8
}
