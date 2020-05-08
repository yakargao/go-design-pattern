#### 单例模式

- 保证一个类仅有一个实例，并提供一个访问它的全局访问点，主要解决一个全局使用的类频繁地创建与销毁。判断系统是否已经有这个单例，如果有则返回，如果没有则创建

- 场景：主要一个对象的场景下，比如1、要求生产唯一序列号，2、EB 中的计数器，不用每次刷新都在数据库里加一次，用单例先缓存起来，3、建的一个对象需要消耗的资源过多，比如 I/O 与数据库的连接等

  ```go
  import "sync"
  
  //Singleton 是单例模式类
  type Singleton struct{}
  
  var singleton *Singleton
  var once sync.Once
  
  //GetInstance 用于获取单例模式对象
  func GetInstance() *Singleton {
  	once.Do(func() {
  		singleton = &Singleton{}
  	})
  
  	return singleton
  }
  ```

  

