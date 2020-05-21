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