/**
* @Author: CiachoG
* @Date: 2020/5/7 20:00
* @Description：
 */
package Creational

import "fmt"

type Shape interface {
	Draw()
}

// 类1
type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("draw a circle")
}

// 类2
type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("draw a Rectangle")
}

// 类2
type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("draw a Square")
}
