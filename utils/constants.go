package utils

type GeneralData byte

const (
	GeneralUser = GeneralData(iota)
	GeneralPassword
)

type NewCmdData byte

const (
	NewArgForEntity = NewCmdData(iota)
	NewFlagPassType
	NewFlagPassLength
)
