# GOlang API Server with Hexagolang architecture
Sample REST API build using echo server.

The code implementation was inspired by port and adapter pattern or known as [hexagonal](blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example):
- **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of all the repository needed and the service itself will be put here.
- **Core**<br/>Contains model/entity and all pure function helper relate to this will
- **Modules**<br/>Contains implementation of interfaces that defined at the business (also called as adapters in hexagonal's term) and also http handler (controller).

# Data initialization

To describe about how port and adapter interaction (separation concerned), this example will have two databases supported. There are MySQL and MongoDB.

MongoDB will become a default databaese in this example. If you want to change into MySQL, update the configuration inside 
[config.yaml](https://raw.githubusercontent.com/hudabikhoir/go-hexagon/master/config/config.yaml) file.

# How To Run Server
Just execute code below in your console
```console
./run.sh
```
