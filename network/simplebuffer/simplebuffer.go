package simplebuffer

import (
	"encoding/binary"
	"unsafe"
)

const MAX_SIZE = 8192

type SimpleBuffer struct {
	data   []byte
	size   int
	offset int
	endian string
}

//endian = "bigEndian" || "litterEndian"
func NewSimpleBuffer(byteorder string) *SimpleBuffer {
	return &SimpleBuffer{data: make([]byte, MAX_SIZE), size: MAX_SIZE, endian: byteorder}
}

func NewSimpleBufferBySize(byteorder string,size int) *SimpleBuffer {
	return &SimpleBuffer{data: make([]byte, size), size: size, endian: byteorder}
}

func NewSimpleBufferByBytes(d []byte, byteorder string) *SimpleBuffer {
	return &SimpleBuffer{data: d, size: len(d), offset: len(d), endian: byteorder}
}

func (this *SimpleBuffer) Data() []byte {
	return this.data[0:this.offset]
}

func (this *SimpleBuffer) Size() int {
	return this.offset
}

func (this *SimpleBuffer) WriteUInt8(i uint8) *SimpleBuffer {
	tsize := int(unsafe.Sizeof(i))
	buf := make([]byte, tsize)
	buf[0] = i
	this.WriteData(buf)
	//this.offset += tsize
	return this
}

func (this *SimpleBuffer) WriteUInt16(i uint16) *SimpleBuffer {
	tsize := int(unsafe.Sizeof(i))
	buf := make([]byte, tsize)
	binary.BigEndian.PutUint16(buf, i)
	this.WriteData(buf)
	//this.offset += tsize
	return this
}

func (this *SimpleBuffer) WriteUInt32(i uint32) *SimpleBuffer {
	tsize := int(unsafe.Sizeof(i))
	buf := make([]byte, tsize)
	binary.BigEndian.PutUint32(buf, i)
	this.WriteData(buf)
	//this.offset += tsize
	return this
}

func (this *SimpleBuffer) WriteUInt64(i uint64) *SimpleBuffer {
	tsize := int(unsafe.Sizeof(i))
	buf := make([]byte, tsize)
	binary.BigEndian.PutUint64(buf, i)
	this.WriteData(buf)
	//this.offset += tsize
	return this
}

func (this *SimpleBuffer) WriteData(d []byte) *SimpleBuffer {
	size := len(d)
	if this.offset + size > this.Length()  {
		tmpBuff := make([]byte,this.Length() + MAX_SIZE)
		copy(tmpBuff, this.data)
		this.data = tmpBuff
	}
	copy(this.data[this.offset:], d)
	this.offset += size
	return this
}

func (this *SimpleBuffer) Length() int{
	return len(this.data)
}

func (this *SimpleBuffer) ReadUInt8() uint8 {
	var _i uint8
	tsize := int(unsafe.Sizeof(_i))
	ret := this.data[0]
	if this.offset <= 0 {
		return uint8(ret)
	}
	copy(this.data[0:], this.data[tsize:this.offset])
	this.offset -= tsize
	return ret
}

func (this *SimpleBuffer) GetUInt8() uint8 {
	return this.data[0]
}

func (this *SimpleBuffer) ReadUInt16() uint16 {
	var _i uint16
	tsize := int(unsafe.Sizeof(_i))
	ret := binary.BigEndian.Uint16(this.data[0:tsize])
	if this.offset <= 0 {
		return ret
	}
	copy(this.data[0:], this.data[tsize:this.offset])
	this.offset -= tsize
	return ret
}

func (this *SimpleBuffer) GetUInt16() uint16 {
	var _i uint16
	tsize := int(unsafe.Sizeof(_i))
	ret := binary.BigEndian.Uint16(this.data[0:tsize])
	if this.offset <= 0 {
		return ret
	}
	return ret
}

func (this *SimpleBuffer) ReadUInt32() uint32 {
	var _i uint32
	tsize := int(unsafe.Sizeof(_i))
	ret := binary.BigEndian.Uint32(this.data[0:tsize])
	if this.offset <= 0 {
		return ret
	}
	copy(this.data[0:], this.data[tsize:this.offset])
	this.offset -= tsize
	return ret
}


func (this *SimpleBuffer) GetUInt32() uint32 {
	var _i uint32
	tsize := int(unsafe.Sizeof(_i))
	ret := binary.BigEndian.Uint32(this.data[0:tsize])
	if this.offset <= 0 {
		return ret
	}
	return ret
}


func (this *SimpleBuffer) ReadUInt64() uint64 {
	var _i uint64
	tsize := int(unsafe.Sizeof(_i))
	ret := binary.BigEndian.Uint64(this.data[0:tsize])
	if this.offset <= 0 {
		return ret
	}
	copy(this.data[0:], this.data[tsize:this.offset])
	this.offset -= tsize
	return ret
}

func (this *SimpleBuffer) GetUInt64() uint64 {
	var _i uint64
	tsize := int(unsafe.Sizeof(_i))
	ret := binary.BigEndian.Uint64(this.data[0:tsize])
	if this.offset <= 0 {
		return ret
	}
	return ret
}

func (this *SimpleBuffer) ReadData(size int) []byte {
	ret := make([]byte, size)
	copy(ret, this.data[0:size])
	if this.offset <= 0 {
		return ret
	}
	copy(this.data[0:], this.data[size:this.offset])
	this.offset -= size
	return ret
}

