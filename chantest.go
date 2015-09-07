package main



import (
	"strings"
	"time"
	"os"
	"fmt"
	"log"
	"runtime"
	"path/filepath"
)


func main () {
	fmt.Printf("logger test \n")
	log.SetOutput(os.Stdout)
	log.Println("my name is log")
	log := NewLogger()
	log.SetOutputDir(`c:\`)
	log.Init("log")
	for i:=0;i<99999;i++ {
		log.LogDebug("i am the logger %v %d\n","hahahaha",i)	
		time.Sleep(time.Millisecond * 100)
	}
	log.Close()
	fmt.Printf("asdasd %v \n","123")
	
}

func Test(format string,v ...interface{}) {
	
}

func Debug(format string,v ...interface{}) {
	
}

func Info(format string,v ...interface{}) {
	
}

func Error(format string,v ...interface{}) {
	
}

func Critial(format string,v ...interface{}) {
	
}

type Logger struct {
	outputDir string
	consoleLv int
	logLv int
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
			fmt.Printf("ERROR %v ",err)
			hFile,err = os.Create(logFullName)
			fmt.Printf("CreateFile err %v \n",err)
			offset = 0
			break
		} else { /* file is exist and go on using it */
			if fi.Size() >= 1024*1024*30 /*30m*/ {
				hFile,_ = os.Create(logFullName)
				offset = 0
				break
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
				if this.offset > 1024*1024*30 {
					this.file.Close()
					file,offset := this.InitFile(this.logName)
					this.file = file
					this.offset = offset
				}
				break	
			case  <-this.chExit :
				this.file.Close()
				return
		}
	}
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
	_, file, line, _ := runtime.Caller(1)
	_,filename := filepath.Split(file)
	log.Printf("=%v:%v:\n",filename,line)
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

func (this *Logger) LogDebug(format string , v ...interface{}) {
	// flush data
	logText := this.formatLogText(format,v...)
	this.chLog <- logText
	
}


func (this *Logger) Close() {
	this.chExit <- true
}

