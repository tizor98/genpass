package utils

type GeneralData byte

const (
	GeneralUser = GeneralData(iota)
)

type NewCmdData byte

const (
	NewArgForEntity = NewCmdData(iota)
	NewFlagPassType
	NewFlagPassLength
)
