package hash


import (
	"hash/crc32"
)



func HashC32(data []byte) uint32{
	const IEEE = 0xedb88320
	var IEEETable = crc32.MakeTable(IEEE)
	return crc32.Checksum(data, IEEETable) 
}