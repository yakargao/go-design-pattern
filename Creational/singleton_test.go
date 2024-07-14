package Creational

import "testing"

func TestGetInstance(t *testing.T) {
	for i := 0; i < 650; i++ {
		go GetInstanceV2()
	}
}
