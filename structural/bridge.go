package structural

import "fmt"

//type message interface {
//	send()
//}
//
//type commonMessageSMS struct {
//
//}
//
//func (c *commonMessageSMS) send(){
//	fmt.Println("sms send")
//}
//
//type commonMessageEmail struct {
//
//}
//
//func (c *commonMessageEmail) send(){
//	fmt.Println("email send")
//}
//
//type UrgencyMessage interface {
//	message
//	watch()
//}
//
//type urgencyMessageSMS struct {
//
//}
//
//func (u *urgencyMessageSMS) watch(){
//	fmt.Println("urgency sms watch")
//}
//
//func (u *urgencyMessageSMS) send(){
//	fmt.Println("urgency sms send")
//}
//
//type urgencyMessageEmail struct {
//
//}
//
//func (u *urgencyMessageEmail) watch(){
//	fmt.Println("urgency email watch")
//}
//
//func (u *urgencyMessageEmail) send(){
//	fmt.Println("urgency email send")
//}
//
//func test(m message){
//	m.send()
//}


// 桥接模式的实现
type messageImpl interface {
	send()
}

type MessageSms struct {

}

func (m *MessageSms) send(){
	fmt.Println("message sms send")
}

type MessageEmail struct {

}

func (m *MessageEmail) send(){
	fmt.Println("message email send")
}

type abstractMessage interface {
	messageImpl
	sendMessage()
}

type commonMessage struct {
	abstractMessage
}

func (c *commonMessage) sendMessage(){

}

type urgencyMessage struct {
	abstractMessage
}

func (u *urgencyMessage) sendMessage(){

}

func (u *urgencyMessage) watch(){

}

func test(){
	m := MessageSms{}
	c := commonMessage{}


}