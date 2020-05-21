/**
* @Author: CiachoG
* @Date: 2020/5/21 16:52
* @Description：
 */
package main

import "fmt"


//Component
type Shape interface {
	Draw()
}

//ConcreteComponent
type Rectangle  struct {

}
func (r *Rectangle )Draw()  {
	fmt.Println("draw a rectangle")
}
//ConcreteComponent
type Circle  struct {
	
}
func (c *Circle )Draw()  {
	fmt.Println("draw a circle")
}

type ShapeDecorator interface {
	Shape
	DrawWithColor()
}
type RedShapeDecorator  struct {
	Shape
}

func (r *RedShapeDecorator)DrawWithColor()  {
	
}

