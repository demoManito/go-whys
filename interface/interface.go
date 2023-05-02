package _interface

type Phone interface {
	Call()
	SendMessage()
}

// NokiaPhone Nokia Phone
type NokiaPhone struct{}

func (NokiaPhone) Call() {
	println("I am Nokia, I can call you!")
}

func (NokiaPhone) SendMessage() {
	println("I am Nokia, I can send message to you!")
}

type MobilePhone interface {
	Phone
	Internet()
}

// XiaoMi MI Phone
type XiaoMi struct {
	NokiaPhone
}

func (XiaoMi) Call() {
	println("I am XiaoMi, I can call you!")
}

func (XiaoMi) Internet() {
	println("I am XiaoMi, I can connect to the Internet!")
}

// Topci1 .
// A. XiaoMi 未实现 MobilePhone 接口
// B. I am XiaoMi, I can call you! \
//    I am Nokia, I can call you! \
//    I am XiaoMi, I can connect to the Internet!
// C. I am XiaoMi, I can call you! \
//    I am Nokia, I can send message to you! \
//    I am XiaoMi, I can connect to the Internet!
// D. I am Nokia, I can call you! \
//    I am Nokia, I can send message to you! \
func Topci1() {
	var phone MobilePhone = &XiaoMi{}
	phone.Call()
	phone.SendMessage()
	phone.Internet()
}
