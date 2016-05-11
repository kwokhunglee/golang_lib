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
