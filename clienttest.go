package main

import (
	"runtime"
	"time"
	"fmt"
	_"gorouter/network/protocol"
	"gorouter/network/simplebuffer"
	"net"
	"gorouter/util"
	_"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_"net/http"
	_"io/ioutil"

)

type RouterInfo struct{
	RemoteAddr string
	SendBytes int64
	RecvBytes int64
	Timestamp int64
}


var (
	RouterConnection [8]RouterInfo
)

func main() {
	
	/*
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	conn := "oa_local" + ":" + `f*(&Dssdsa)s` + "@tcp(" + "127.0.0.1:51816" + ")/" + "db_oa_enterprise" + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)
	orm.SetMaxIdleConns("default", 3)
	orm.SetMaxOpenConns("default", 3)
	
	resp,err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Printf("http Get Error %v \n",err.Error())
	}
	b,_ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Get Data %v \n",string(b))
	
	return
	*/
	
	for i:=0;i<100;i++ {
		go newTcpConn()
		time.Sleep( 10 * time.Second)
	}
	newTcpConn()
}

func newTcpConn() {
	defer util.TraceCrashStack()
	fmt.Print("Go Router Client Runing \n")
	connClient, err := net.Dial("tcp", "127.0.0.1:9093")
	if err != nil {
		fmt.Print("Client Connecting error \n")
		return
	}

	buffer := simplebuffer.NewSimpleBuffer("bigEndian")
	buffer.WriteUInt32(0x3) //len
	buffer.WriteUInt8(0x80)
	buffer.WriteUInt8(1)
	buffer.WriteData([]byte("0")) 
	fmt.Printf("send data %v \n", buffer.Data())
	connClient.Write(buffer.Data())
	
/*
	sb := simplebuffer.NewSimpleBuffer("bigEndian")
	chBuffer := make(chan *simplebuffer.SimpleBuffer)
	
	exitChan := make(chan bool)
	go unpack(chBuffer,exitChan,connClient)
	
	for {
		fmt.Print("Recv Looping  \n")
		buf := make([]byte, 4096)
		n, err := connClient.Read(buf)
		if err != nil {
			fmt.Printf("Client Read Buffer Failed %v %v\r\n", err, n)
			break
		}
		fmt.Printf("\n\nrecv data %v \n\n",buf[0:n])
		sb.WriteData(buf[0:n])
		chBuffer <- sb
		fmt.Printf("numbers of goroutine %v \n",runtime.NumGoroutine())
	}
	
	fmt.Printf("recv loop exit \n")
*/
}

/*pack unpack*/
func unpack(chanBuffer chan *simplebuffer.SimpleBuffer,exitChan chan bool,conn net.Conn) {
	defer util.TraceCrashStack()
	for {
		recvBuf,_ := <- chanBuffer
		for {
			if recvBuf.Size()  <= 0 {
				break;
			}
			
			/* create a temp buffer */
			tmpData := make([]byte,recvBuf.Size())
			copy(tmpData,recvBuf.Data())
			tmpBuffer := simplebuffer.NewSimpleBufferByBytes(tmpData,"bigEndian")
			totalSize := int(tmpBuffer.Size())
			bodySize := int(tmpBuffer.ReadUInt32())
			if totalSize - 4 /*4 is header size*/ < bodySize {
				/*  the pack is not full , go on recv */
				continue
			}
			/*when parse enough space .. */	
			
			bSize := recvBuf.ReadUInt32()
			var mouID int
			var cmdID int
			if bodySize > 2 {
				fmt.Printf("big body %v \n",bodySize)
				mouID = int(recvBuf.ReadUInt8())
				cmdID = int(recvBuf.ReadUInt8())
			} else {
				fmt.Printf("little body %v \n",bodySize)
				recvBuf.ReadData(bodySize)
			}
			fmt.Printf("bSize %v mid %v cid %v \n",bSize,mouID,cmdID)
			
			if mouID != 0x80 || cmdID != 0x01 {
				continue
			}
			/*  handle return data
			*   proto [  <<header/32>>
						<<module/8,command8>>
						<<size/8>> 
							Loop<<1,IP1,IP2,IP3,IP4,Port:16,RecvBytes:64,SendBytes:64>> ] 
			*/
			fmt.Printf("Router State: \n")
			routerConns := int(recvBuf.ReadUInt8())
			for i:=0;i<routerConns;i++ {
				succFlag := int(recvBuf.ReadUInt8())
				if succFlag == 0 {
					continue
				}
				ip1 := recvBuf.ReadUInt8()
				ip2 := recvBuf.ReadUInt8()
				ip3 := recvBuf.ReadUInt8()
				ip4 := recvBuf.ReadUInt8()
				port := recvBuf.ReadUInt16()
				recvBytes := recvBuf.ReadUInt64()
				sendBytes := recvBuf.ReadUInt64()
				fmt.Printf("  %v.%v.%v.%v:%v recv: %v send: %v \n",ip1,ip2,ip3,ip4,port,recvBytes,sendBytes)
				RouterConnection[i].RemoteAddr = fmt.Sprintf("%v.%v.%v.%v:%v",ip1,ip2,ip3,ip4,port)
				RouterConnection[i].SendBytes = int64(sendBytes)
				RouterConnection[i].RecvBytes = int64(recvBytes)
				RouterConnection[i].Timestamp = time.Now().Unix()
			}
			fmt.Printf("exit chan trigger \n")
			conn.Close()
			runtime.Goexit()
		}
	}
}