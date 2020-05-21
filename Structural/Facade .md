### 外观模式

- 外观模式又称门面模式，使用一个外观类对象对复杂的对象进行访问，屏蔽复杂对象的细节。

- 两个角色

  - System系统类 
  - Facade外观类

  ![Facade](../assets/Facade.jpg)

- 实现

  ```
  /**
  * @Author: CiachoG
  * @Date: 2020/5/21 20:29
  * @Description：
   */
  package main
  
  import "fmt"
  
  type SystemA struct {
  
  }
  
  func (s *SystemA)doSomething()  {
  	fmt.Println("a")
  }
  type SystemB struct {
  
  }
  
  func (s *SystemB)doSomething()  {
  	fmt.Println("b")
  }
  type SystemC struct {
  
  }
  
  func (s *SystemC)doSomething()  {
  	fmt.Println("c")
  }
  
  
  type Facade struct {
  	A SystemA
  	B SystemB
  	C SystemC
  }
  
  func (f *Facade)Invoke()  {
  	f.A.doSomething()
  	f.B.doSomething()
  	f.C.doSomething()
  }
  ```

  

