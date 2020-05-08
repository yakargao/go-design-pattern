#### 创建者模式

- 使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式。

- 适合一个对象的零部件，但是构造算法（过程）不变的场景

- 建造者由4个部分组成：

  - 产品（Production）
  - 抽象建造者（Builder）
  - 具体建造者（Concrete Builder）
  - 指挥者（Director）：建造过程所在

- 例子:

- 代码:汽车装配过程

  ```go
  /**
  * @Author: CiachoG
  * @Date: 2020/5/8 10:00
  * @Description：
   */
  package main
  
  import "fmt"
  
  //Production
  type Car struct {
  	name string
  }
  
  
  //Builder
  type Builder interface {
  	newProduct()
  	//方法私有，因为简单的一个步骤并不能获得完成的car
  	buildWheels()Builder
  	buildEngine()Builder
  	buildBody()Builder
  	getResult()interface{}
  }
  
  //ConcreteBuilder
  type CarBuilder struct {
  	car *Car
  }
  
  func (c *CarBuilder)newProduct(){
  	c.car = &Car{
  		name: "宝马",
  	}
  }
  func (c *CarBuilder)buildWheels()Builder{
  	fmt.Println("装配完轮胎")
  	return c
  }
  func (c *CarBuilder)buildEngine()Builder{
  	fmt.Println("装配完引擎")
  	return c
  }
  func (c *CarBuilder)buildBody()Builder{
  	fmt.Println("装配完整体")
  	return c
  }
  func (c *CarBuilder)getResult()interface{}{
  	return c.car
  }
  
  //Director
  type Director struct {
  	builder Builder
  }
  
  func (d *Director)SetBuilder(builder Builder){
  	d.builder = builder
  }
  
  func (d *Director)Construct()*Car {
  	fmt.Println("开始装配……")
  	d.builder.newProduct()
  	return d.builder.buildBody().buildEngine().buildWheels().getResult().(*Car)
  }
  
  func main()  {
  	carBuilder :=  new(CarBuilder)
  	carDirector := new(Director)
  	carDirector.SetBuilder(carBuilder)
  	car := carDirector.Construct()
  	fmt.Println(car)
  }
  
  ```

  