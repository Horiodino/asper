package vm
//
//import (
//	"fmt"
//
//	"github.com/horiodino/asper/internal/logger"
//)
//
//type UserContext struct {
//	Username string
//	Password string
//}
//
//type asper struct {
//	Create Creator
//	Delete Delete
//	Get    Get
//	Modify Modify
//	List   ListOfResources
//	// Help()
//	// VersionVM()
//	// ConfigVM() *UserContext
//	// SSHVM(name string) *UserContext
//	MetricsVM MetricsVM
//	// Attach    Attach
//	// CheckHealthVM(id string) *UserContext
//}
//
//func (c *asper) Help() {
//	fmt.Println("Running Help")
//}
//
//type Cli struct {
//	isadmin bool
//}
//
//func NewLocalClient() *asper {
//	fmt.Println("Running NewLocalClient")
//	return &asper{
//		// Create:    &create{},
//		MetricsVM: &metrice{},
//	}
//}
//
//type InboundRule struct {
//	iprangeFrom logger.Asperstring
//	iprangeTo   logger.Asperstring
//	protocol    logger.Asperstring
//	ports       logger.Asperstring
//}
//
//type OutboundRule struct {
//}
//
//func AsperClient() *asper {
//	fmt.Println("Running AsperClient")
//	return &asper{
//		Create:    &create{},
//		MetricsVM: &metrice{},
//	}
//}
//
//func main() {
//
//	metrice := &metrice{UserID: "1", IsUSer: true, InstanceID: "1"}
//	asper := &asper{MetricsVM: metrice}
//	val := asper.Modify.InitilizeModifyClient("")
//	res := val.ModifyVM("")
//	res.AddFirewallRule()
//
//}
