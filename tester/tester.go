package tester

import (
	"os"
	"time"
)

func Run(ex string) {
	for {
		time.Sleep(100 * time.Millisecond)
		bs := []byte(time.Now().Format("2006-01-02 15:04:05.999999999") + "\r\n")
		f, err := os.OpenFile(ex+".data", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		if _, err := f.Write(bs); err != nil {
			panic(err)
		}
		f.Close()
	}
}
