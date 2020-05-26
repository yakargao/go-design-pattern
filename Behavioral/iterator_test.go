/**
* @Author: CiachoG
* @Date: 2020/5/26 20:06
* @Description：
 */
package Behavioral

import (
	"testing"
)

func TestLangIterator(t *testing.T) {
	iterator := NewLangRepo().GetIterator()
	for iterator.HasNext() {
		t.Error(iterator.Next())
	}
}
