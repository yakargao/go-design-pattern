#### 桥接模式

- 对于一个系统或者一个对象有多个维度的变化情况，如果采用多层继承的方式，一个维度的变化就会影响另外其他维度。桥接模式正是针对这种情况，抽象化与实现化解耦，使得二者可以独立变化

- 桥接模式有4个角色

  - 抽象类
  - 扩展抽象类
  - 实现类
  - 具体实现类

- 实现

  ```go
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
  ```

  

