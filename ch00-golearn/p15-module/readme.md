go里的module相当于java的maven里的pom的概念

go mod init github.com/rechard/modules_test 执行后会在项目里建立一个go.mod，这个类似pom.xml

然后比如我们要用到github.com/pquerna/otp 来生成一次性密码，就在那个目录下使用
go get -u github.com/pquerna/otp,接下来就可以直接import使用了
