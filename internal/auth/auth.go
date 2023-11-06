package auth


// this will implement all the auth related function like login, logout, checkrole, etc

type Auth interface {
	Login() (bool, error)
	Logout() (bool, error)
	IsLoggedIn() (bool, error)
	CheckRole() (bool, error)
}


// TODO we will authenticate the user from the database in case of the remote context
// we will create the db for the first time and we will store the token in the config file
// once the db is created we will use the token from the config file and we will authenticate the user

// for the cli context point of view we will use the token from the config file and we will authenticate the user
// but we dont use any database for the cli context we use simple .config file for the cli context to save the token
// and all other information

type AuthImpl struct {
	// TODO

}

func (auth *AuthImpl) Login() (bool, error) {

	return true, nil
}

func (auth *AuthImpl) Logout() (bool, error) {

	return true, nil
}

func (auth *AuthImpl) IsLoggedIn() (bool, error) {

	return true, nil
}

func (auth *AuthImpl) CheckRole() (bool, error) {

	return true, nil
}

func InitializeAuth(auth *AuthImpl) (*AuthImpl, error) {

	return &AuthImpl{}, nil
}