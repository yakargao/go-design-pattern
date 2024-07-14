/**
* @Author: CiachoG
* @Date: 2020/5/8 10:41
* @Description：
 */
package Creational

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	carBuilder := new(CarBuilder)
	carDirector := new(Director)
	carDirector.SetBuilder(carBuilder)
	car := carDirector.Construct()
	t.Log(car)
}
