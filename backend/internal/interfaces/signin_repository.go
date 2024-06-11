package interfaces

type SigninRepository interface {
	Signin(username string, password string) (string, error)
}
