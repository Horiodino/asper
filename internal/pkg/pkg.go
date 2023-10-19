package pkg

import "fmt"

// this will cointail all implementataion like createvm, deletevm, etc all

func (cli *Cli) CreateVM(name string, vmos string, cpu int, memory int, disk int) (error) {
	fmt.Println("Running CreateVM")
	return nil
}

func (cli *Cli) DeleteVM(name string) (error) {
	return nil
}

func (cli *Cli) GetVM(name string) (error) {
	return nil
}

func (cli *Cli) UpdateVM(name string, vmos string, cpu int, memory int, disk int) (error) {
	return nil
}

func (cli *Cli) ListVM() (error) {
	return nil
}

func (cli *Cli) HelpVM() (error) {
	return nil
}

func (cli *Cli) VersionVM() (error) {
	return nil
}

func (cli *Cli) ConfigVM() (error) {
	return nil
}

func (cli *Cli) SSHVM(name string) (error) {
	return nil
}

func (cli *Cli) MetricsVM(name string) (error) {
	return nil
}

func (cli *Cli) IsLoggedIn() (error) {

	return nil
}

func (cli *Cli) checkrole() (error) {

	return nil
}

func (cli *Cli) Login() (error) {

	return nil
}

func (cli *Cli) Logout() (error) {

	return nil
}
