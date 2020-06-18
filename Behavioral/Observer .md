#### 观察者模式

- 定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。golang中典型的发布订阅者模式

- docker源码中的发布订阅模式：

  ```
  
  ​```
  /**
  * @Author: CiachoG
  * @Description：
  * @Date: 2020/4/20 17:18
   */
  package pub_sub
  
  import (
  	"sync"
  	"time"
  )
  
  type (
  	subscriber chan interface{}//订阅者
  	topicFunc  func(v interface{}) bool //主题过滤器
  )
  
  type Publisher struct {
  	mutex       sync.RWMutex //读写锁
  	buffer      int          //订阅队列的缓存大小
  	timeout     time.Duration
  	subscribers map[subscriber]topicFunc
  }
  
  //新增发布者
  func NewPublisher(timeout time.Duration, buffer int) *Publisher {
  	return &Publisher{
  		buffer:      buffer,
  		timeout:     timeout,
  		subscribers: make(map[subscriber]topicFunc),
  	}
  }
  
  //订阅全部主题
  func (p *Publisher) Subscribe() subscriber {
  	return p.SubscribeTopic(nil)
  }
  
  //订阅一个主题
  func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
  	ch := make(subscriber, p.buffer)
  	p.mutex.Lock()
  	p.subscribers[ch] = topic
  	p.mutex.Unlock()
  	return ch
  }
  
  //退出订阅
  func (p *Publisher) Evict(sub subscriber) {
  	p.mutex.Lock()
  	delete(p.subscribers, sub)
  	p.mutex.Unlock()
  }
  
  //发布一个主题内容
  func (p *Publisher) Publish(v interface{}) {
  	p.mutex.RLock()
  	defer p.mutex.RUnlock()
  	var wg sync.WaitGroup
  	for sub, topic := range p.subscribers {
  		wg.Add(1)
  		p.sendTopic(sub, topic, v, &wg)
  	}
  	wg.Wait()
  }
  
  //发送主题内容
  func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
  	defer wg.Done()
  	if topic != nil && !topic(v) {
  		return
  	}
  	select {
  	case sub <- v:
  	case <-time.After(p.timeout):
  	}
  }
  
  //关闭发布对象
  func (p *Publisher) Close() {
  	p.mutex.Lock()
  	defer p.mutex.Unlock()
  	for sub := range p.subscribers {
  		delete(p.subscribers, sub)
  		close(sub)
  	}
  }
  
  ​```
  
  ```

  