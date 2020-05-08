#### 抽象工厂模式

- 提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。主要解决接口选择的问题。

- ==使用场景==：系统的产品有多于一个的产品族，而系统只消费其中某一族的产品。比如生成不同操作系统的程序

- 关系

  ![abstractFactory](../assets/abstractFactory.jpg)

- go 实现

  

  ``` go
  //------------------Shape------------------------
  //Shape接口
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
  //----------------------------Shape--------
  ```

  ```go
  //------Color
  type Color interface {
  	Fill()
  }
  
  //类1
  type Red  struct {
  
  }
  
  func (r *Red)Fill()  {
  	fmt.Println("red")
  }
  
  //类2
  type Green  struct {
  
  }
  
  func (g *Green)Fill()  {
  	fmt.Println("Green")
  }
  
  //类2
  type Blue struct {
  
  }
  
  func (b *Blue)Fill()  {
  	fmt.Println("Blue")
  }
  ```

  

  ```go
  //抽象工厂接口
  type AbstractFactory interface {
  	GetColor(typ int)Color
  	GetShape(typ int)Shape
  }
  //实现工厂接口的两个工厂类
  //工厂类
  type ColorFactory struct {
  
  }
  
  func (c *ColorFactory)GetColor(typ int)Color{
  	switch typ {
  	case 1:
  		return &Red{}
  	case 2:
  		return &Green{}
  	case 3:
  		return &Blue{}
  	}
  	return nil
  }
  
  func (c *ColorFactory)GetShape(typ int)Shape{
  	return nil
  }
  
  //工厂类
  type ShapeFactory struct {
  
  }
  
  func (s *ShapeFactory)GetShape(typ int)Shape{
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
  
  
  func(s *ShapeFactory)GetColor(typ int)Color{
  	return nil
  }
  
  type FactoryProducer struct {
  
  }
  //工厂类生成器
  func (f *FactoryProducer)getFactory(typ int)AbstractFactory{
  	switch typ {
  	case 1:
  		return &ShapeFactory{}
  	case 2:
  		return &ColorFactory{}
  	}
  	
  }
  ```

  

  - 



