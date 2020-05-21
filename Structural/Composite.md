#### 组合模式

- 组合模式是通过一个对象代表一组相同类型的对象，对这组对象的操作仅需通过这个对象，类似于现实生活中一个leader代表一个团队的所有成员

- 实现

  ```
  /**
  * @Author: CiachoG
  * @Date: 2020/5/21 16:01
  * @Description：
   */
  package main
  
  type Student struct {
  	Name string
  	Age int
  	Classmates []Student
  }
  
  func (s *Student)add(student Student){
  	s.Classmates = append(s.Classmates,student)
  }
  func (s *Student)remove(student Student){
  	//…………
  }
  func (s *Student)List()[]Student  {
  	return s.Classmates
  }
  ```

  

  