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


