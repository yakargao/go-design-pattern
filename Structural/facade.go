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

