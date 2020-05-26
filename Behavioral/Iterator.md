### 迭代器模式

- 提供用于顺序访问集合对象，而不用暴露集合对象的底层表示

- 由于迭代器模式将存储数据和遍历数据的职责分离，增加新的聚合类需要对应增加新的迭代器类，类的个数成对增加，这在一定程度上增加了系统的复杂性。

- 角色：

  - Iterator 接口
  - Iterator 实现类
  - Container 容器类
  - Container 实现类

- 实现：

  ````
  /**
  * @Author: CiachoG
  * @Date: 2020/5/26 19:38
  * @Description：
   */
  package Behavioral
  
  type Iterator  interface {
  	HasNext()bool
  	Next()interface{}
  }
  type Container interface {
  	GetIterator()*Iterator
  }
  
  type CallBack func(int)interface{}
  
  type LangIterator struct {
  	index int
  	len int
  	cb CallBack
  }
  
  func (l *LangIterator)HasNext()bool{
  	return l.index < l.len
  }
  func (l *LangIterator)Next()interface{} {
  	next := l.cb(l.index)
  	l.index ++
  	return next
  }
  
  type LangRepository  struct {
  	Langs []string
  }
  
  func (r *LangRepository)GetIterator()*LangIterator  {
  	return &LangIterator{index: 0,len: len(r.Langs),cb: r.Next}
  }
  
  func NewLangRepo()*LangRepository{
  	return &LangRepository{
  		Langs: []string{"golang","java","python"},
  	}
  }
  func (r *LangRepository)Next(idx int)interface{}{
  	return r.Langs[idx]
  }
  ````

  