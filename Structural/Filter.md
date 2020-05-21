#### 过滤器模式

- 过滤器是设计模式里面比较简单的模式，使用不同的过滤标准对一组对象进行过滤

- 三个角色

  - 过滤器接口
  - 过滤器实现类
  - 对象

- 实现

  ```go
  /**
  * @Author: CiachoG
* @Date: 2020/5/21 11:27
  * @Description：
   */
  package main
  
  
  type Person struct {
  	name string
  	age int
  }
  
  
  //过滤接口
  type Criteria interface {
  	meetCriteria([]Person)[]Person
  }
  
  //过滤实现1
  type JuvenileCriteria struct {
  	
  }
  
  func (j *JuvenileCriteria) meetCriteria(persons []Person)[]Person  {
  	juveniles := make([]Person,0)
  	for _,p := range persons{
  		if p.age < 18 {
  			juveniles = append(juveniles,p)
  		}
  	}
  	return  juveniles
  }
  
  //过滤实现2
  type MiddleAgedCriteria struct {
  
  }
  
  func (m *MiddleAgedCriteria) meetCriteria(persons []Person)[]Person  {
  	mas := make([]Person,0)
  	for _,p := range persons{
  		if p.age < 18 {
  			mas = append(mas,p)
  		}
  	}
  	return  mas
  }
  
  //过滤实现2
  type SeniorsCriteria struct {
  
  }
  
  func (s *SeniorsCriteria) meetCriteria(persons []Person)[]Person  {
  	seniors := make([]Person,0)
  	for _,p := range persons{
  		if p.age < 18 {
  			seniors = append(seniors,p)
  		}
  	}
  	return  seniors
  }
  ```
  
  

