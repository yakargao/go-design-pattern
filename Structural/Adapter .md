#### 适配器模式

- 现有的接口需要转化为客户类期望的接口，这样保证了对现有类的重用

- 在适配器模式中需要定义一个适配器（adapter）来完成适配的工作，类似于电源适配器或者编码解码器的作用。如果没有这个适配器，那么适配的工作就需要在目标类（Target）中完成，这样待适配类（Adaptee）的功能对于调用它的目标类来说就不透明了。适配工作也可以在待适配类（Adaptee），这样就需要对其进行改动，新增接口

- 思想类似的地方：在前后端分离的项目中，服务端可能与不用的客户端采用不同的协议进行交互，如果在control直接根据不同协议处理，会使得control层变得疲惫不堪。正确的做法是加入协议层，把协议相关（可能是thrift，http或者grpc）的数据转化成协议无关的数据，交给control处理。

- 适配器模式包含4个角色：

  - Target：目标类
  - Adapter：适配器类
  - Adaptee：适配者类
  - Client:客户端类

- 实现

  ```go
  /**
  * @Author: CiachoG
* @Date: 2020/5/9 11:53
  * @Description：
   */
  package Structural
  
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
  
  ```
  
  ```go
  //测试
  func TestAdapter(t *testing.T)  {
  	adaptee := NewAdaptee()
  	adapter := NewAdapter(adaptee)
  	t.Log(adapter.Request())
  }
  ```
  
  