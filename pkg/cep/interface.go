package cep

type CEP interface {
	FindLocation(cep string) (string, error)
}
