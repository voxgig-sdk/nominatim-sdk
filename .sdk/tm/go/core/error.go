package core

type NominatimError struct {
	IsNominatimError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewNominatimError(code string, msg string, ctx *Context) *NominatimError {
	return &NominatimError{
		IsNominatimError: true,
		Sdk:              "Nominatim",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *NominatimError) Error() string {
	return e.Msg
}
