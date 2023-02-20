package main

type myint int

type IPerson interface {
	GetType() string
}

type Person struct {
	Name string
	Age  int
	Type string
}

//不要写成这种
/*func (this Person) Walk() {

}*/

func (this *Person) Walk() {
	println(this.Name, " walking")
}

func (this *Person) GetType() string {
	return this.Type
}

// 相当于继承
type SuperMan struct {
	Person
}

func (this *SuperMan) Fly() {
	println("superman=", this.Name, " flying")
}

func info(p IPerson) {
	println(p.GetType())
}

// java里的Object,如要区分的话，可以使用instanceof
// go里的interface{}是万能断言,类似java里的object,
// 如何区分底层的数据到底是哪种类型,使用arg.(string)
func info2(arg interface{}) {
	s, ok := arg.(string)
	if ok {
		println(s)
	}
}
func main() {
	//var i myint
	//println(i)
	//fmt.Printf("%t", i)

	person := Person{Name: "rechard", Age: 20}
	person.Walk()
	superMan := SuperMan{
		person,
	}
	superMan.Fly()

	var pi IPerson
	pi = &person
	info(pi)

	pi = &superMan
	info(pi)
}
