package logger

/*
	these are the interface for the logger
	// it will be used to log the error and other information
*/

/*

	teh logger be called globally , so like we create interface it is asper , asper is the our project name
	so we will create the interface as asper.String() , asper.Bool() .....

*/

type Logger interface {
	// string will convert any type to string

	String([]string) (string, error)
	Bool([]string) (bool, error)
	Int([]string) (int, error)
	Float([]string) (float64, error)
	Bytes([]string) ([]byte, error)
	JSON([]string) (interface{}, error)
	JSONLines([]string) ([]interface{}, error)
	Table(map[string]string, []map[string]string) error
	Error(error) error
	Invalid(err error)

	/*
		this will be used to log the error it will find what the error type if its a custom error then it will print the custom error else it will print the error
		 or invalid err type
	*/

}

type LoggerImpl struct {
}

func NewLogger() *LoggerImpl {
	return &LoggerImpl{}
}

func (l *LoggerImpl) String([]string) (string, error) {
	panic("implement me")
}

func (l *LoggerImpl) Bool([]string) (bool, error) {
	panic("implement me")
}

func (l *LoggerImpl) Int([]string) (int, error) {
	panic("implement me")
}

func (l *LoggerImpl) Float([]string) (float64, error) {
	panic("implement me")
}

func (l *LoggerImpl) Bytes([]string) ([]byte, error) {
	panic("implement me")
}

func (l *LoggerImpl) JSON([]string) (interface{}, error) {
	panic("implement me")
}

func (l *LoggerImpl) JSONLines([]string) ([]interface{}, error) {
	panic("implement me")
}

func (l *LoggerImpl) Table(map[string]string, []map[string]string) error {
	panic("implement me")
}

func (l *LoggerImpl) Error(error) error {
	panic("implement me")
}

func (l *LoggerImpl) Invalid(err error) {
	panic("implement me")
}
