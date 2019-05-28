## Polymorphism

Polymorphism is the essence of object-oriented programming: the ability to treat objects of different types uniformly as long as they adhere to the same interface. Go interfaces provide this capability in a very direct and intuitive way.

Polymorphism in Go is achieved with the help of interfaces. Interfaces can be implicitly implemented in Go. A type implements an interface if it provides definitions for all the methods declared in the interface. To demonstrate this I've created 
a simple finance example for available payment methods and its processing.


Payment module is consumed by Pay module, there is also a client package that shows the instantiation of the objects and use of each method for payment. 