package protocol


type IPCProtocol struct {
	ModuleId int
	CommandId int
	Data interface{}
	RetChan interface{}
}