package utils

type GeneralFlags byte

const (
	GeneralUser = GeneralFlags(iota)
)

type NewFlags byte

const (
	NewFlagForEntity = NewFlags(iota)
)
