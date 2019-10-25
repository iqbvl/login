package Login

type LoginUsecase interface {
	Login(username string, password string) (token string)
}
