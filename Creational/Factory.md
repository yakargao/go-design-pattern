#### 工厂模式

- 工厂模式使其创建过程延迟到子类进行,屏蔽复杂的对象的生成，主要解决接口选择的问题

- **使用场景：** 1、日志记录器：记录可能记录到本地硬盘、系统事件、远程服务器等，用户可以选择记录日志到什么地方。 2、数据库访问，当用户不知道最后系统采用哪一类数据库，以及数据库可能有变化时。 3、设计一个连接服务器的框架，需要三个协议，"POP3"、"IMAP"、"HTTP"，可以把这三个作为产品类，共同实现一个接口。

- 关系

  ![factory](../assets/factory.jpg)

- go 实现

  ```go
  /**
  * @Author: CiachoG
  * @Date: 2020/5/7 20:00
  * @Description：
   */
  package Behavioral
  
  import "fmt"
  
  type Shape interface {
  	Draw()
  }
  
  //类1
  type Circle struct {
  
  }
  
  func (c *Circle)Draw()  {
  	fmt.Println("draw a circle")
  }
  
  //类2
  type Rectangle struct {
  
  }
  
  func (r *Rectangle)Draw()  {
  	fmt.Println("draw a Rectangle")
  }
  
  //类2
  type Square struct {
  
  }
  
  func (s *Square)Draw()  {
  	fmt.Println("draw a Square")
  }
  
  //工厂类
  type ShapeFactory struct {
  
  }
  
  func (s *ShapeFactory)GetSharpe(typ int)Shape{
  	switch typ {
  	case 1:
  		return &Circle{}
  	case 2:
  		return &Rectangle{}
  	case 3:
  		return &Square{}
  	}
  	return nil
  }
  
  
  ```

  

