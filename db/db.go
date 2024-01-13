package db

type database struct {
	Token            string `json:"token"`
	Context          string `json:"context"`
	OperationType    string `json:"operationType"`
	OperationTime    string `json:"operationTime"`
	OperationStatus  string `json:"operationStatus"`
	OperationMessage string `json:"operationMessage"`
}

// db that is going to be used is mongodb

// this func is going to be used to write the log in the database takes the data and the option of any type
func (d *database) WriteLog(msg string, option interface{}) error {
	return nil
}

// this func is going to be used to write the error in the database
// or we just return the error directly
func (d *database) WriteError(msg string, option interface{}) error {
	return nil
}

// the data that is going to be written in the database we dont know what it is
// so we use interface{} to accept any type of data,  so by that based on the data
// we can know what type of data it is and we can write it in the database
// if its creating vm make a tabel like
// ----------------------------
// | vmName | vmId | vmStatus |
// ----------------------------
// |  vm1   |  1   |  running |
// ----------------------------
// |  vm2   |  2   |  running |
// ----------------------------

// we can tell the db to write what ever data we want in the database
//based on the cases of datatype that we have in the interface{}

// this function is going to be used to write the data in the database takes the data and the option of any type
func (d *database) WriteInfo(msg string, option interface{}) error {
	return nil
}
