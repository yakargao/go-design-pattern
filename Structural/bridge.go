/**
* @Author: CiachoG
* @Date: 2020/5/9 16:37
* @Description：抽象类：笔的大小 实现类：笔的颜色
 */
package main

import "fmt"

type Pen interface {
	ColorImpl
	SetColorImpl(impl ColorImpl)
}

//--------实现类----------
type ColorImpl interface {
	AddColor()
}

type RedColorImpl struct {

}

func (r *RedColorImpl)AddColor()  {
	fmt.Println("add red")
}

type BlueColorImpl struct {

}
func (b *BlueColorImpl)AddColor()  {
	fmt.Println("add blue")
}

type GreenColorImpl struct {

}
func (g *GreenColorImpl)AddColor()  {
	fmt.Println("add green")
}
//--------------------


//细化抽象类:大
type LargePen struct {
	ColorImpl
}

func (l *LargePen)SetColorImpl(impl ColorImpl)  {
	l.ColorImpl = impl
}
func (l *LargePen)AddSize()  {
	fmt.Println("add Large")
}
func (l *LargePen)GenPen(){
	l.AddColor()
	l.AddSize()
	fmt.Println("generate a pen")
}

//细化抽象类:中
type MediumPen struct {
	ColorImpl
}

func (m *MediumPen)SetColorImpl(impl ColorImpl)  {
	m.ColorImpl = impl
}
func (m *MediumPen)AddSize()  {
	fmt.Println("add MediumPen")
}
func (m *MediumPen)GenPen(){
	m.AddColor()
	m.AddSize()
	fmt.Println("generate a pen")
}

//细化抽象类:中
type SmallPen struct {
	ColorImpl
}

func (s *SmallPen)SetColorImpl(impl ColorImpl)  {
	s.ColorImpl = impl
}
func (s *SmallPen)AddSize()  {
	fmt.Println("add MediumPen")
}
func (s *SmallPen)GenPen(){
	s.AddSize()
	s.AddColor()
	fmt.Println("generate a pen")
}

func main(){
	largeRedPen := new(LargePen)
	largeRedPen.SetColorImpl(new(RedColorImpl))
	largeRedPen.GenPen()
}