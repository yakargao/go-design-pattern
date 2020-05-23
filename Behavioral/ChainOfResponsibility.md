#### 责任链模式

- 为请求对象创建一条处理链，逐级处理
- 两个角色：
  - 抽象处理接口
  - 具体处理类
- 典型运用:  - 消息过滤器 - 权限拦截器 - log打印…… 
- 实现

```
第一个实例可以参考gin中间件的设计
```

```
/**
* @Author: CiachoG
* @Date: 2020/5/23 10:33
* @Description：
 */
package Behavioral

import "log"

type Handler interface {
	Handle(context string)
	Next(handler Handler,context string)
}

type LogHandler struct {
	NextHandler Handler
}

func (l *LogHandler)Handle(context string)  {
	//do the responsibility of logger
	log.Println("打印日志")
	context = context + "已打印完日志"
	l.Next(l.NextHandler,context)
}
func (l *LogHandler)Next(handler Handler,context string){
	handler.Handle(context)
}

type FilterHandler struct {
	NextHandler Handler
}

func (f *FilterHandler)Handle(context string)  {
	//do the responsibility of FilterHandler
	f.Next(f.NextHandler,context)
}
func (f *FilterHandler)Next(handler Handler,context string){
	handler.Handle(context)
}
//.....
```

