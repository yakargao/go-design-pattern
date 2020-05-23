#### 装饰器模式

- 装饰器模式用于对现有对象的功能的扩展，类似于继承的效果，但是继承是静态的，用户不能控制扩展的方式和时机。

- 包含的角色有四种

  - Component: 抽象构件

  - ConcreteComponent: 具体构件

  - Decorator: 抽象装饰类

  - ConcreteDecorator: 具体装饰类

    ![Decorator](../assets/Decorator.jpg)

- 实现

  ```
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
  
  ```

