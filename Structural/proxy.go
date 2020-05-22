/**
* @Author: CiachoG
* @Date: 2020/5/22 21:12
* @Description：
 */
package main

import "fmt"

type Subject interface {
	 request()
}
type RealSubject struct {

}

func (r *RealSubject)request()  {
	fmt.Println("request")
}

type Proxy struct {
	Subject
}

func NewProxy()Proxy{
	return Proxy{Subject:&RealSubject{}}
}

func (p *Proxy)Do()  {
	p.beforeRequest()
	p.request()
	p.afterRequest()
}
func (p *Proxy)beforeRequest(){

}
func (p *Proxy)afterRequest(){

}




