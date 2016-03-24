package simplebuffer

import (
	"fmt"

	"testing"

)


/*
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
*/

func Test_Scale(t *testing.T) {
	buffer := NewSimpleBufferBySize("BigEndian",4096)
	fmt.Printf("buffer %v \n",buffer)
	
	b := make([]byte,10000)
	for i:=0;i<10;i++ {
		buffer.WriteData([]byte("123"))
		buffer.WriteUInt8(1)
	
		buffer.WriteUInt16(12)
	
		buffer.WriteUInt32(32)
		buffer.WriteUInt64(64)
		buffer.WriteData([]byte("987654321"))
		buffer.WriteData(b)
		buffer.Data()
	}
	

	//fmt.Printf("buffer %v \n",buffer.Data())
	/*
	for i:=0;i<1000;i++ {
		x1 := buffer.ReadData(3)
		x2 := buffer.ReadUInt8()
		x3 := buffer.ReadUInt16()
		x4 := buffer.ReadUInt32()
		x5 := buffer.ReadUInt64()
		x6 := buffer.ReadData(9)
		fmt.Printf("x1 :%v\nx1 :%v\nx1 :%v\nx1 :%v\nx1 :%v\nResult : %v\n",x1,x2,x3,x4,x5,x6)
	}

	*/
	
	
}
