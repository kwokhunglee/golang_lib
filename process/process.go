/*
线程管理
kwokhung lee
guoxiongli@alog.cc
2016-04-03

此包的作用是在启动时在指定文件夹产生PID文件，
在PID文件中写入当前程序的PID，
并在程序退出时将这个文件删除。

通常用于服务启动判断程序是否已经启动
*/

package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Process struct {
	pidfile string
	pid     int
}

/*
使用方法如下：
package main

import (
	"bufio"
	"log"
	"os"
	"alog.cc/lib/process"
)

func main() {
	var p, err = process.InitProcess("/run/test/4434")
	if err != nil {
		log.Println(err)
		p.CloseProcess()
		return
	}
	defer p.CloseProcess()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

}
*/
func InitProcess(pidfile string) (p *Process, err error) {
	p = &Process{}
	p.pidfile = pidfile
	p.pid = os.Getpid()
	log.Println("Program Running, PID Is:", p.pid)

	dir := filepath.Dir(pidfile)
	if err = os.Chdir(dir); err != nil {
		if err = os.Mkdir(dir, os.ModePerm); err != nil {
			return
		}
		return
	}

	if err = ioutil.WriteFile(p.pidfile, []byte(fmt.Sprintf("%d\n")), 0644); err != nil {
		return
	}
	return
}

func (p *Process) CloseProcess() (err error) {
	log.Println("Program Stop, PID Is:", p.pid)
	err = os.Remove(p.pidfile)
	return err
}
