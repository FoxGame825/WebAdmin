package libs

import (
	"os"
	"time"
	"strings"
)

//打印内容到文件中
//tracefile(fmt.Sprintf("receive:%s",v))
func MyLog(str_content string)  {
	fd,_:=os.OpenFile("MyLog.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====\n",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

func MyLogByFileName( fileName string, str_content string)  {
	fd,_:=os.OpenFile(fileName + ".txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====\n",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

