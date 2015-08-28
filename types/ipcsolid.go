package types


type IPCSolid struct {
	ModuleID int
	CommandID int
	Data interface{}
	RetChan interface{}
}
