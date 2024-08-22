package componentes

type MesiFlags = uint8

const (
	M MesiFlags = iota // 0
	E
	S
	I
)

type Tags = uint8

const ( //read hit, read miss...
	RH Tags = iota //0
	RM
	WH
	WM
)
