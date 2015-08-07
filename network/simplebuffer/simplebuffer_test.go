package simplebuffer

import (

	"testing"

)



func Test_SimpleBuffer(t *testing.T) {
	for i:= 0;i<1;i++ { //主要性能消耗在内存分配上
		_buffer := NewSimpleBuffer("bigEndian")
	//for i:= 0;i<1;i++ { 
		_buffer.WriteUInt16(5)
		//fmt.Printf("buffer %v \r\n", _buffer.Data())
		_buffer.WriteUInt16(6)
		_buffer.WriteUInt32(7)
		_buffer.WriteUInt64(8)
		_buffer.WriteUInt16(9)
		_buffer.WriteData([]byte("string"))
		//fmt.Printf("buffer %v \r\n", _buffer.Data())
		_buffer.ReadUInt16()
		_buffer.ReadData(2)
		_buffer.ReadUInt32()
		_buffer.ReadUInt64()
		_buffer.ReadUInt16()
		_buffer.ReadData(6)
		
		//fmt.Printf("read %v buffer %v \r\n", _buffer.ReadUInt16(), _buffer.Data())

		//fmt.Printf("read %v buffer %v \r\n", _buffer.ReadData(2), _buffer.Data())
		//fmt.Printf("read %v buffer %v \r\n", _buffer.ReadUInt32(), _buffer.Data())
		//fmt.Printf("read %v buffer %v \r\n", _buffer.ReadUInt64(), _buffer.Data())
		//fmt.Printf("read %v buffer %v \r\n", _buffer.ReadUInt16(), _buffer.Data())

		//fmt.Printf("read %v buffer %v \r\n", string(_buffer.ReadData(6)), _buffer.Data())
	}
}


