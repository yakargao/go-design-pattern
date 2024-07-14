package main

import "go-design-pattern/Creational"

func main() {
	for i := 0; i < 30; i++ {
		go Creational.GetInstanceV2()
	}
}
