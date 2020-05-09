/**
* @Author: CiachoG
* @Date: 2020/5/9 11:53
* @Description：
 */
package main

//Target
type Target interface {
	Request() string  //client需要的方法
}

//----------------Adaptee-------------------------------//

//Adaptee
type Adaptee interface {
	SpecificRequest() string //原来的方法
}

//NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//AdapteeImpl
type adapteeImpl struct{}

//SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}


//-------------------------Adapter-----------------------------//

//Adapter
type adapter struct {
	Adaptee
}

//Request
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
//NewAdapter
func NewAdapter(adaptee Adaptee)*adapter  {
	return &adapter{
		Adaptee: adaptee,
	}
}






