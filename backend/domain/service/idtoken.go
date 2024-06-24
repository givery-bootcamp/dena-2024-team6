package service

type IdtokenService interface {
	Generate(id string) (string, error)
	Verify()
}
