# 项目章节
- ch01 基础的server
   包含的功能有
   
   1.开启服务器
   

- ch02




# go语言的总结

### go并发

goroutine 类似于线程，但是可以根据需要创建多个 goroutine 并发工作。
goroutine 是由 Go 语言的运行时调度完成，而线程是由操作系统调度完成
goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务，它是Go语言并发设计的核心

goroutine 在多核 cpu 环境下是并行的，如果代码块在多个 goroutine 中执行，那么我们就实现了代码的并行。

使用普通函数创建 goroutine
格式 ： go 函数名( 参数列表 )

```go
func main() {
    go func() {
 		fmt.Println("in go func ")
 	}()
 
 	fmt.Println("in go main ")
 
 	select {}
}
```

所有 goroutine 在 main() 函数结束时会一同结束。


协程/线程
协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，
这些用户级线程的调度也是自己实现的。

线程：一个线程上可以跑多个协程，协程是轻量级的线程。

上面的不是很理解，协程讲的不就是java的线程？

channel
channel 是Go语言在语言级别提供的 goroutine 间的通信方式。
我们可以使用 channel 在两个或多个 goroutine 之间传递消息。

对比java 中管道用于2个流之间的数据传输

channel 是类型相关的，也就是说，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定

```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

Go语言既然以并发编程作为语言的最核心优势

### 命名规则

- 包名称
 
  保持package的名字和目录保持一致，包名应该为小写单词，不要使用下划线或者混合大小写

- 文件命名
  应该为小写单词，使用下划线分隔各个单词
  
- 结构体命名
  采用驼峰命名法，首字母根据访问控制大写或者小写  
  
- 接口命名
   文件为
  
 ### Zinx的GoPath
 
 ~~GOPATH是一个开发环境目录的意思，下面必须包含bin、pkg、src，然后再src下面新建项目zinx
 project setting 里的gopath设置成 D:\dev-code\my_github_code\go~~
 
 告别GOPATH，快速使用 go mod（Golang包管理工具）
 go mod 类似java的maven
 使用go mod 管理项目，就不需要非得把项目放到GOPATH指定目录下，你可以在你磁盘的任何位置新建一个项目。
  
 
 步骤
 1. 新建一个名为 wserver 的项目，项目路径 D:\test\wserver 
 2. 进入项目目录 D:\test\wserver 里，新建一个 go源码文件： main.go
 3. go mod init wserver 
 
 注意：
 GOPROXY 需要设置，默认的从https://proxy.golang.org/下载
 包下载路径是在
 C:\Users\sdrag\go\pkg\mod\cache\download\
 目前还不知道如何更改
 
 
 idea的project setting->go->go modules里
 enviroments 点击后加入GOPROXY ,  值填入 https://proxy.golang.org/
 填写完毕后需要重启
 
 
 
 
### 字节数组
```go
 //直接字符串转成字节数组
 []byte("i love you")
``` 
 
 
 ## 问题记录
 1. ch01 里的server 如果发生了错误
 ```go
         addr,err:=net.ResolveTCPAddr(self.TCPVersion,fmt.Sprintf("%s:%d",self.IP,self.Port))
		if err!=nil{
            //这里发生了错误
			fmt.Println("resolve tcp addr error:",err)
            //但是照样会打印出start Zinx server success，因为用了go func(){}()处理
            //该怎么改？
            return 
		}
 ```
  