package types


type IPCSolid struct {
	ModuleId int
	CommandId int
	Data interface{}
	RetChan interface{}
}