package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
	"strings"
	"time"
)

func main() {
	c, err := ftp.Dial("188.131.241.21:21", ftp.DialWithTimeout(5*time.Second))

	fmt.Println("I'm here ----1111")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("I'm here ----1111")


	err = c.Login("ftptest", "ftp123456")

	fmt.Println("I'm here ----2222")

	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn
	fmt.Println(c.CurrentDir())
	list,_ := c.NameList("/home/ftptest")
	c.Stor("/home/ftptest/test.txt", strings.NewReader("asdad"))

	fmt.Println(list)




	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
