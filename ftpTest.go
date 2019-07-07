package main
import (
	"fmt"
	"os"
	ftp4go "github.com/shenshouer/ftp4go"
)
func main() {
	ftpClient := ftp4go.NewFTP(0) // 1 for debugging
	//connect
	_, err := ftpClient.Connect("188.131.241.21", 22, "")
	if err != nil {
		fmt.Println("The connection failed")
		fmt.Println(err)
		os.Exit(1)
	}
	defer ftpClient.Quit()
	_, err = ftpClient.Login("root", "Anseek2019", "")
	if err != nil {
		fmt.Println("The login failed")
		os.Exit(1)
	}
	//Print the current working directory
	var cwd string
	cwd, err = ftpClient.Pwd()
	if err != nil {
		fmt.Println("The Pwd command failed")
		os.Exit(1)
	}
	fmt.Println("The current folder is", cwd)
}