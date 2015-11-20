package protocol

import (
	"errors"
	"gorouter/network/simplebuffer"
)

const PROTOCOL_HEADER int = 4
const PROTOCOL_MODULE_ID int = 1
const PROTOCOL_COMMAND int = 1

type Protocol struct {
	Header   int
	ModuleId uint8
	Command  uint8
	Data     []byte
}

func NewProtocal() *Protocol {
	return &Protocol{}
}

func (this *Protocol) ToBytes() []byte {
	buffer := simplebuffer.NewSimpleBuffer("bigEndian")

	buffer.WriteUInt32(uint32(this.Header))
	buffer.WriteUInt8(uint8(this.ModuleId))
	buffer.WriteUInt8(uint8(this.Command))
	buffer.WriteData(this.Data)
	return buffer.Data()
}

func (this *Protocol) ParseFromParam(mod uint8, cmd uint8, data []byte) *Protocol {
	this.Header = PROTOCOL_MODULE_ID + PROTOCOL_COMMAND + len(data)
	this.ModuleId = mod
	this.Command = cmd
	this.Data = data
	return this
}

func (this *Protocol) PraseFromData(data []byte, size int) (*Protocol, error) {
	if size < PROTOCOL_HEADER+PROTOCOL_MODULE_ID+PROTOCOL_COMMAND {
		return nil, errors.New("Buffer size is too small to protocal size")
	}

	//don't use the src point
	//make a copy with the input data
	dataCopy := make([]byte,size)
	copy(dataCopy,data)
	
	buffer := simplebuffer.NewSimpleBufferByBytes(dataCopy, "bigEndian")
	this.Header = int(buffer.ReadUInt32())
	if this.Header + PROTOCOL_HEADER != size {
		return nil, errors.New("Protocal header size error \n")
	}
	
	this.ModuleId = (buffer.ReadUInt8())	
	this.Command = (buffer.ReadUInt8())
	this.Data = buffer.ReadData(this.Header - ( PROTOCOL_MODULE_ID + PROTOCOL_COMMAND))
	
	return this, nil
}
