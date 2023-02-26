package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	ServerIP          string
	ServerPort        int
	Conn              net.Conn
	Name              string
	mode              int
	RequestPendingMap map[string]chan string
}

func (c *Client) Menu() int {
	var mode int

	fmt.Println("0.退出")
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改用户名")

	fmt.Scanln(&mode)

	if mode >= 0 && mode <= 3 {
		return mode
	}

	return -1
}

// 更名模式
func (c *Client) UpdateName() {
	println(">>>>请输入你的新名字: ")
	var newName string
	_, err := fmt.Scanln(&newName)

	if newName == "" {
		println("不能输入空白字符")
		return
	}

	if err != nil {
		println("输入有误:", err)
		return
	}

	//封装成更名的消息格式
	updateNameMsg := "rename|" + newName + "\n"
	c.Conn.Write([]byte(updateNameMsg))
}

// 公聊模式
func (c *Client) PublicChat() {
	println(">>>>你已进入公聊模式，请输入内容, exit表示退出")
	var content string

	fmt.Scanln(&content)

	for content != "exit" {
		if len(content) != 0 {
			_, err := c.Conn.Write([]byte(content + "\n"))
			if err != nil {
				println("write content error", err)
				break
			}
		}

		content = ""
		println(">>>>你已进入公聊模式，请输入内容, exist表示退出")
		fmt.Scanln(&content)
	}

	println("你已退出公聊模式")

}

func (c *Client) OnlineUsers() {

	//输出聊天对象
	c.Conn.Write([]byte("who"))
}

// 私聊模式
func (c *Client) PrivateChat() {
	//聊天对象
	var chatPersonName string = ""
	for chatPersonName != "exit" {
		println(">>>>你已进入私聊模式，请选择聊天对象:")

		c.OnlineUsers()
		fmt.Scanln(&chatPersonName)

		existChan := c.CheckUserExist(chatPersonName)
		resp := <-existChan
		if resp == "0" {
			println(chatPersonName + "不存在，请重新选择")
			chatPersonName = ""
			continue
		}

		println(">>>>你现在可以和" + chatPersonName + "聊天")
		var content string

		fmt.Scanln(&content)
		for content != "exit" {
			c.Conn.Write([]byte("to|" + chatPersonName + "|" + content + "\n"))

			println(">>>>你现在可以和" + chatPersonName + "聊天")
			fmt.Scanln(&content)

		}

		println("你已退出和" + chatPersonName + "私聊模式")
		chatPersonName = ""
	}
	println("你已退出私聊模式")
}

func (c *Client) Run() {

	for choseItem := -1; choseItem != 0; {
		choseItem = c.Menu()
		switch choseItem {
		case -1:
			println("输入无效，请输入0到3")
			break
		case 0:
			println("你已退出")
			return
		case 1:
			//println("当前是公聊模式")
			c.PublicChat()
			break
		case 2:
			//println("当前是私聊模式")
			c.PrivateChat()
			break
		case 3:
			c.UpdateName()
			break
		}
	}
}

func (c *Client) DealResponse() {
	bytes := make([]byte, 4096)
	for {
		n, err := c.Conn.Read(bytes)
		if err != nil {
			println("error in read ", err)
		}
		str := string(bytes[:n])

		if index := strings.Index(str, "|"); index != -1 {
			strArr := strings.Split(str, "|")
			reqNum := strArr[0]
			res := str[1]
			c.RequestPendingMap[reqNum] <- string(res)
		} else {
			println(str)
		}
	}

	//io.Copy(os.Stdout, c.Conn)
}

func (c *Client) CheckUserExist(name string) chan string {
	//1.生成一个系列号
	unique_number := string(time.Now().Nanosecond())
	//2.发送请将请求放到一个地方存起来
	resp := make(chan string)
	c.RequestPendingMap[unique_number] = resp
	req := fmt.Sprintf("user|%d|%s", unique_number, name)
	c.Conn.Write([]byte(req))

	return resp
}

func NewClient(serverIp string, serverPort int) *Client {

	client := &Client{
		ServerIP:          serverIp,
		ServerPort:        serverPort,
		RequestPendingMap: make(map[string]chan string),
	}

	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		println("error in dial", err)
		return nil
	}

	go client.DealResponse()

	client.Conn = conn
	return client
}

var serverIP string
var serverPort int

func init() {
	flag.StringVar(&serverIP, "s", "127.0.0.1", "identify serverIP, default to 127.0.0.1")
	flag.IntVar(&serverPort, "p", 9999, "identify serverPort,default to 9999")
}

func main() {
	flag.Parse()
	client := NewClient(serverIP, serverPort)
	if client != nil {
		println("connect to server success")
	} else {
		println("connect to server fail")
	}

	client.Run()

}
