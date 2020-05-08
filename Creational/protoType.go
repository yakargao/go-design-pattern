/**
* @Author: CiachoG
* @Date: 2020/5/7 21:34
* @Description：
 */
package main

type Cloneable  interface {
	Clone()Cloneable
}


type PrototypeManager struct {
	prototypes map[string]Cloneable //可克隆对象集中管理
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}