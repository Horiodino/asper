package db

// there will be two interface first of all a simpler interface that will have just acesss to the database
// just for checking not for modifying the database

// and second interface will be for modifying the database
// it will be only called when we have some deeeep operation to perform on the database
// like a request came to server to createa a new vm
// so first we will check the user that requested to create vm
// then we authenticate the user bu the first interface and all other things necessary

// then when all things done , when some validation of user input is done
// then we will call the second interface to create the vm and save hte information in the database

// bu that way we will have two interfaces
// one for checking and one for modifying the database

// todo implement hash to generate toke as hash and only save the hash in the database
//  by doing that we will have more security
type ExternalLevel interface {
	IsuserExist(username string) (bool, error)
	IsuserValid(username string, password string) (bool, error)
	LastLogin(username string) (string, error)
}

type InternalLevel interface {
	WriteLog(msg string, option interface{}) error
	WriteError(msg string, option interface{}) error
	WriteInfo(msg string, option interface{}) error
}
