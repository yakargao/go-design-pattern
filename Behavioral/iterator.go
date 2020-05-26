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