package common

type Pair[F any, S any] struct {
	First  F
	Second S
}

func CreatePair[F any, S any](first F, sencond S) Pair[F, S] {
	return Pair[F, S]{first, sencond}
}
