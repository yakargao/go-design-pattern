
package main
/**
* @Author: CiachoG
* @Date: 2020/5/7 20:00
* @Description：
 */

import "fmt"



type AbstractFactory interface {
	GetColor(typ int)Color
	GetShape(typ int)Shape
}

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
	fmt.Println("draw a Rectangle")
}

//类2
type Blue struct {

}

func (b *Blue)Fill()  {
	fmt.Println("draw a Square")
}

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

func (f *FactoryProducer)getFactory(typ int)AbstractFactory{
	switch typ {
	case 1:
		return &ShapeFactory{}
	case 2:
		return &ColorFactory{}
	}
	
}
