package protocol

import (
	"errors"
	"github.com/TriangleGo/gorouter/network/simplebuffer"
)

const PR_HEAD int = 4
const PR_MID int = 1
const PR_CID int = 1

type Protocol struct {
	Header   int
	ModuleId uint8
	Command  uint8
	Data     []byte
}

func NewProtocal() *Protocol {
	return &Protocol{}
}

func NewProtocalByParams(m,c uint8,d []byte) *Protocol {
	return &Protocol{Header:len(d) + PR_MID + PR_CID, ModuleId:m,Command:c,Data:d}
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
	this.Header = PR_MID + PR_CID + len(data)
	this.ModuleId = mod
	this.Command = cmd
	this.Data = data
	return this
}

func (this *Protocol) PraseFromData(data []byte, size int) (*Protocol, error) {
	if size < PR_HEAD + PR_MID+PR_CID {
		return nil, errors.New("Buffer size is too small to protocal size")
	}

	//don't use the src point
	//make a copy with the input data
	dataCopy := make([]byte,size)
	copy(dataCopy,data)
	
	buffer := simplebuffer.NewSimpleBufferByBytes(dataCopy, "bigEndian")
	this.Header = int(buffer.ReadUInt32())
	if this.Header + PR_HEAD != size {
		return nil, errors.New("Protocal header size error \n")
	}
	
	this.ModuleId = (buffer.ReadUInt8())	
	this.Command = (buffer.ReadUInt8())
	this.Data = buffer.ReadData(this.Header - ( PR_MID + PR_CID))
	
	return this, nil
}

func (this *Protocol) GetModuleId() uint8 {
	return this.ModuleId
}

func (this *Protocol) GetCommandId() uint8 {
	return this.Command
}
