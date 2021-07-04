package base

import "errors"

var (
	DBQueryError = errors.New("db query error")

	DBQueryNotFound = errors.New("not found segment")

	DBExecError = errors.New("db exec sql error")

	DBUpdateNoAffect = errors.New("db update affect zero rows")

	IdOverPresetMaximum = errors.New("id exceeds the preset maximum")
)
