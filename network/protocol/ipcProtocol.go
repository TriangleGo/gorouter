package protocol


type IPCProtocol struct {
	ModuleId string
	CommandId string
	Data interface{}
	RetChan interface{}
}