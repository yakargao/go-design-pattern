原型模式

- ==实现了一个原型接口，该接口用于创建当前对象的克隆==。当直接创建对象的代价比较大时，则采用这种模式。例如，一个对象需要在一个高代价的数据库操作之后被创建。我们可以缓存该对象，在下一个请求时返回它的克隆，在需要的时候更新数据库，以此来减少数据库调用。

- ==**使用场景：**==  1、类初始化需要消化非常多的资源，这个资源包括数据、硬件资源等。 2、性能和安全要求的场景。3、通过 new 产生一个对象需要非常繁琐的数据准备或访问权限，则可以使用原型模式。 4、一个对象多个修改者的场景：一个对象需要提供给其他对象访问，而且各个调用者可能都需要修改其值时，可以考虑使用原型模式拷贝多个对象供调用者使用。 5、在实际项目中，原型模式很少单独出现，一般是和==工厂方法模式==一起出现，通过 clone 的方法创建一个对象，然后由工厂方法提供给调用者。

- go 实现

  ```go
  //可以克隆的对象必须实现的接口
  type Cloneable  interface {
  	Clone()Cloneable
  }
  type PrototypeManager struct {
  	prototypes map[string]Cloneable //可克隆对象集中管理
  }
  
  func NewPrototypeManager() *PrototypeManager {
  	return &PrototypeManager{
  		prototypes: make(map[string]Cloneable),
  	}
  }
  
  func (p *PrototypeManager) Get(name string) Cloneable {
  	return p.prototypes[name]
  }
  
  func (p *PrototypeManager) Set(name string, prototype Cloneable) {
  	p.prototypes[name] = prototype
  }
  ```

  

