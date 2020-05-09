/**
* @Author: CiachoG
* @Date: 2020/5/9 13:57
* @Description：
 */
package main

import "testing"

func TestAdapter(t *testing.T)  {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	t.Log(adapter.Request())
}
