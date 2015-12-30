package logger


import (
	"strings"
	"time"
	"os"
	"fmt"
	"runtime"
	"path/filepath"
	_"github.com/TriangleGo/gorouter/lib/goconfig"
)

const (
	FILE_SIZE = 1024*1024*5 /* 1m for total byte for unit*/
)


type Logger struct {
	outputDir string
	consoleLv int
	logConsolLv int
	logFileLv int
	file *os.File
	offset int64
	chLog chan string
	chExit chan bool
	logName string
}

func NewLogger() *Logger{
	return &Logger{chLog: make(chan string)	 ,
			chExit:make(chan bool)}
}

func (this *Logger) Init(logName string ) {
	this.logName = logName
	file,offset := this.InitFile(logName)
	this.file = file
	this.offset = offset
	go this.LogProc()
}


func (this *Logger) InitFile(logName string) (*os.File,int64) {
	Y,M,D := time.Now().Date()
	logNameWithDate := fmt.Sprintf("%v_%v_%v_%v",logName,Y,int(M),D)
	logExt := `.txt`
	
	var fi os.FileInfo
	var hFile *os.File
	var offset int64
	var err error
	
	for i:=0;i<100;i++ {
		logFullName := fmt.Sprintf("%s%s_%d%s",this.outputDir,logNameWithDate,i,logExt)
		fi , err = os.Lstat(logFullName)
		if err != nil { /* when the file is not exist then create it */
			//fmt.Printf("ERROR %v ",err)
			hFile,err = os.Create(logFullName)
			//fmt.Printf("CreateFile err %v \n",err)
			offset = 0
			break
		} else { /* file is exist and go on using it */
			if fi.Size() >= FILE_SIZE  {
				continue
			} else  {
				hFile,_ = os.OpenFile(logFullName,os.O_WRONLY,os.ModeAppend)
				offset = fi.Size() 
				break
			}
		} //end of else
	} // end for
	return hFile,offset
}

func (this *Logger) LogProc() {
	for {
		select {
		case data,_ := <-this.chLog :
			// return when can't get way to create file 
			if this.file == nil {
				return 
			}
			//flush to file
			n,err := this.file.WriteAt([]byte(data),this.offset);
			// add offset and size 
			if err != nil {
				
			} else {
				this.offset += int64(n)
			}
			// split file /*bigger than 30m*/
			if this.offset > FILE_SIZE {
				this.file.Close()
				file,offset := this.InitFile(this.logName)
				this.file = file
				this.offset = offset
			}
			break	
		case  <-this.chExit :
			this.file.Close()
			return
		} // end select
	} //end for
}

func (this *Logger) SetLogLevel(cmdLv int,fileLv int) {
	this.logConsolLv = cmdLv
	this.logFileLv = fileLv
}

func  (this *Logger) SetOutputDir(dir string) {
	fix := dir[ len(dir) - 1:len(dir) ]  
	
	if fix == `\` || fix == `/`{
		this.outputDir = dir
	} else if  strings.Contains(dir,`\`) {
		this.outputDir = dir + `\`
	} else if  strings.Contains(dir,`/`) {
		this.outputDir = dir + `/`
	} else {
		this.outputDir = dir + `/`
	}
	
}

func (this *Logger)getFileAndLine() (string,int){
	_, file, line, _ := runtime.Caller(5)
	_,filename := filepath.Split(file)
	return filename,line
}

func (this *Logger) getFormat() string {
	file,line := this.getFileAndLine()
	Y,M,D := time.Now().Date()
	h := time.Now().Hour()
	m := time.Now().Minute()
	s := time.Now().Second()
	format := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d %s:%d ",Y,M,D,h,m,s,file,line)
	return format
}

func (this *Logger) formatLogText(format string,v ...interface{}) string {
	formatHeader := this.getFormat()
	logText := fmt.Sprintf( formatHeader + format,v...)
	return logText
}

func (this *Logger) Log(lv int,format string , v ...interface{}) {
	// flush data
	mapLevel := map[int]string{5:"[TEST]",4:"[DEBG]",3:"[INFO]",2:"[EROR]",1:"[CRTL]",0:"[NONE]"}
	logText := this.formatLogText(mapLevel[lv] + " " + format,v...)
	//check if printout the console log
	if lv > this.logConsolLv {
		//don't print out log
	} else {
		fmt.Print(logText)
	}
	
	//check if printout the file log
	if lv > this.logFileLv {
		//don't print out log
	} else {
		this.chLog <- logText	
	}
}


func (this *Logger) Close() {
	this.chExit <- true
}