### 代理模式

- 给某一个对象提供一个代 理，并由代理对象控制对原对象的引用。客户对象对原对象的操作只能通过代理对象进行，类似于跳板机的作用，典型应用是Spring Aop
- 三个角色：
  - Subject: 抽象主题角色
  - Proxy: 代理主题角色
  - RealSubject: 真实主题角色

![../_images/Proxy.jpg](../assets/Proxy.jpg)

- 时序图

  ![../_images/seq_Proxy.jpg](../assets/seq_Proxy.jpg)

- **缺点：**由于在客户端和真实主题之间增加了代理对象，因此 有些类型的代理模式可能会造成请求的处理速度变慢

- **常见代理：**

  - **远程(Remote)代理**：为一个位于不同的地址空间的对象提供一个本地 的代理对象，这个不同的地址空间可以是在同一台主机中，也可是在 另一台主机中，远程代理又叫做大使(Ambassador)。
  - **虚拟(Virtual)代理**：如果需要创建一个资源消耗较大的对象，先创建一个消耗相对较小的对象来表示，真实对象只在需要时才会被真正创建。
  - **Copy-on-Write代理**：它是虚拟代理的一种，把复制（克隆）操作延迟 到只有在客户端真正需要时才执行。一般来说，对象的深克隆是一个 开销较大的操作，Copy-on-Write代理可以让这个操作延迟，只有对象被用到的时候才被克隆。
  - **保护(Protect or Access)代理**：控制对一个对象的访问，可以给不同的用户提供不同级别的使用权限。
  - **缓冲(Cache)代理**：为某一个目标操作的结果提供临时的存储空间，以便多个客户端可以共享这些结果。
  - **防火墙(Firewall)代理**：保护目标不让恶意用户接近。
  - **同步化(Synchronization)代理**：使几个用户能够同时使用一个对象而没有冲突。
  - **智能引用(Smart Reference)代理**：当一个对象被引用时，提供一些额外的操作，如将此对象被调用的次数记录下来等

- 实现：

  ```
  /**
  * @Author: CiachoG
  * @Date: 2020/5/22 21:12
  * @Description：
   */
  package main
  
  import "fmt"
  
  type Subject interface {
  	 request()
  }
  type RealSubject struct {
  
  }
  
  func (r *RealSubject)request()  {
  	fmt.Println("request")
  }
  
  type Proxy struct {
  	Subject
  }
  
  func NewProxy()Proxy{
  	return Proxy{Subject:&RealSubject{}}
  }
  
  func (p *Proxy)Do()  {
  	p.beforeRequest()
  	p.request()
  	p.afterRequest()
  }
  func (p *Proxy)beforeRequest(){
  
  }
  func (p *Proxy)afterRequest(){
  
  }
  ```

  

  

