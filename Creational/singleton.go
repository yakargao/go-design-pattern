/**
* @Author: CiachoG
* @Date: 2020/5/7 21:02
* @Description：
 */
package Creational

import (
	"fmt"
	"sync"
)

// Singleton 是单例模式类
type singleton struct{}

var instance *singleton
var once sync.Once

// GetInstance 用于获取单例模式对象
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

// 并发可能有问题的
func GetInstanceV2() *singleton {
	if instance == nil {
		fmt.Println("create new")
		instance = new(singleton)
		return instance
	}
	fmt.Println("Single instance already created.")
	return instance
}
