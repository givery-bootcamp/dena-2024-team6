package service

type IdtokenService interface {
	Generate(id string) (string, error)
	VerifyIDToken(token string) (int, error)
}
