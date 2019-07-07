package main

import (
	"strings"
	ftp4go "github.com/shenshouer/ftp4go"
	"fmt"
	"flag"
	"github.com/larspensjo/config"
	"time"
)

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)

//配置文件
var TOPIC = make(map[string]string)
var TOPIC1 = make(map[string]string)
//配置文件

func main() {

	/*读取配置*/
	cfg, err1 := config.ReadDefault(*configFile)
	if err1 != nil {
		restart("读取配置文件失败")
	}
	if cfg.HasSection("server") {
		section, err1 := cfg.SectionOptions("server")
		if err1 == nil {
			for _, v := range section {
				options, err1 := cfg.String("server", v)
				if err1 == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	if cfg.HasSection("client") {
		section, err1 := cfg.SectionOptions("client")
		if err1 == nil {
			for _, v := range section {
				options, err1 := cfg.String("client", v)
				if err1 == nil {
					TOPIC1[v] = options
				}
			}
		}
	}
	/*读取配置*/

	/*连接ftp*/
	ftpClient := ftp4go.NewFTP(0)
	_, err := ftpClient.Connect(TOPIC["address"], 22, "")
	if err != nil {
		fmt.Println(TOPIC)
		restart("连接失败")
	}

	fmt.Println(1231313123)
	defer ftpClient.Quit()
	_, err = ftpClient.Login(TOPIC["name"], TOPIC["password"], "")
	if err != nil {
		restart("登陆失败")
	}
	fmt.Println(1231313123)
	/*连接ftp*/

	/*读取目录*/
	var files []string
	c1 := make(chan string)
	go func(){
		if files, err = ftpClient.Nlst(TOPIC["path"]); err != nil {
			restart("读取目录失败")
		}
		c1 <- "读取目录成功"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 10):
		restart("读取目录超时")
	}
	if len(files) == 0 {
		restart("无可下载文件")
	}
	/*读取目录*/

	/*执行下载*/
	for i := 0; i < len(files); i++ {
		ch := files[i]
		str := strings.Replace(ch, " ", ",", -1)
		s := strings.Split(str,",")
		x := len(s)-1
		file1 := s[x]
		if err = ftpClient.DownloadResumeFile(TOPIC["path"]+"/"+file1, TOPIC1["path"]+file1, false); err == nil{
			fmt.Println(TOPIC["path"]+"/"+file1+"--->下载成功")
			_, err = ftpClient.Delete(TOPIC["path"]+"/"+file1)
			if err == nil {
				fmt.Println(TOPIC["path"]+"/"+file1+"--->删除成功")
			}else{
				restart("删除失败")
				break
			}
		}else{
			restart("下载失败")
			break
		}
	}
	/*执行下载*/
	restart("新的一轮执行")
}

func restart(err string){
	fmt.Println("重启--------->"+err)
	time.Sleep(time.Second * 40)
	main()
}