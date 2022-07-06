# 1. SOLID

## 1.1 Single Responsability Pronciple
A type should have only one primary responsabilty
* it should have one reason to change, the reason being related to its prime responsability
* separation of concerns - different types/packages handling different, independent tasks/problems
* an example of violating single responsability princile is a "God object"

## 1.2 Open-Closed Principle
Types should be:
 -  Open for extension
 -  Closed for modification
 
 If it is tested once, the tests should maintain for the functionalities. 

Example: Enterprise pattern of "Specification"

## 1.3 Liskov Substitution Principle
Doesn't apply to Go because go doesn't have inheritance.
"Objects of a superclass should be replaceable with objects of its subclasses without breaking the application"

The behavior of implementers of a particular interface should not brake the core behaviors
that we rely on.

A nice examples of this principle I found @:

https://blog.knoldus.com/what-is-liskov-substitution-principle-lsp-with-real-world-examples/

## 1.4 Interface Segregation Principle
You simply shoulnd't put too much into interface.
YAGNI - You Ain't Going to Need It
Stick to the go std library inerface style ("Reader","Writer","Closer", etc.)
Try to maintain the interfaces small and if they get bigger, just split them.

If we construct an interface with many methods and then we need a type that doesn't
have all the methods it will not pass as an argument to the function that expects the interface. 

## 1.5 Dependency Inversion Principle
High level modules should not depend on low level modules. 
Both of them should depend on abstractions.
By low-level module we mean closer to the hardware: a datastore, communication, system-level stuff. 
High-level would be the business-logic stuff.
Both of them should depend on abstractions: typically we mean interfaces in Go.


## 2. BUILDER
- A builder is a separate component used for building an object
- Some objects are simple and can be created in a single constructor call
- Other objects are more complicated to creae them at once
- Having a factory function with 10 arguments is not productive
- Instead, opt for piecewise (piece-by-piece) constructions
- Builder provides an API for constructing an object step-by-step
- To make builder fluent, rerturn the receiver - it allows chaining actions/methods
- Different facets of an object can be constructed with different builders 
    working in tandem via a common struct
- 

## 2.1 Creational Builder
- Fluent interfaces enable to chain calls together
- it returns the receiver object so you can chain it
- chaining the methods helps building the object

## 2.2 Builder Facets
- It splits the building of an object into multiple logical builder structs
- It allows, through the types, accessing the methods of builders that refer to the single fields
    of the principal struct
- Itallos chaining the methods to set the fields (in any order)

## 2.3 Builder Parameter
- Hide the objects behind the builder and API's user doesn't touch the object directly
- It forces the API client to use builder instead of the object directly
- The builder ensures the object is constructed correctly
- The API client has to provide the action function that initializes email object through the builder 
- The sendMailImpl function uses the email object hidden behind the builder
- We use the fluent interface to construct the email thorough the builder

## 2.4 Functinal Builder
- functional programming with builder pattern in go
- enables delayed application
- the builder instead of just doing the modifications in place can keep a list of actions, 
    or list of changes to perform upon the object that is being constructed
- with the .Build() we create the default implementation of the object and we apply the actions stored on the object


## 3. FACTORY
A component responsible solely for the wholesale (no-piecewise) creation of object.

- a helper function for making struct instances
- any entity that can take care of object creation
- ways of controlling how an object is constructed
- Motivation: object creation becomes too complicated
- If there are a lot of elements to initialize we may want to give some default values to some elements
- Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to:
    * A separate function (Factory Function, a.k.a. Constructor)
    * A separate struct (Factory)

## 4. PROTOTYPE
An initialized (partially or fully) object that you copy (clone) and make use of

- complicated objects aren't designed from scratch
    * they reiterate existing designs
- an existing (partially or fully constructed) design is a prototype
- we make a copy of the prototype and customize it
    * requires a 'deep copy' support (if I copy a struct, what happens to the pointers?
        there might be several pointers pointing to the same object)
- we can make the cloning convenient (via a Factory)

To implement a prototype, partially construct an object and store it somewhere.
Deep copy the prototype.
Customize the resulting instance.
A prototype factory provides a convenient API for using prototypes.

## 5. SINGLETON
Acomponent that is instantiated only once

- for some components it only makes sense to have  one in the system
    * database repository
    * object factory - if the factory is stateless (doesn't have any fields) - 
        there is no point in having more than one instance of the factory
- the construction cost is expensive (loadin a db)
    * we want to do it only once
    * we give everybody the same instance
- we want to prevent anyone creating additional copies
- need to take care of lazy instantiation

1. Lazy one-time initialization with sync.Once
2. Adhere to dependency inversion: depend on interfaces, not concrete types
    (we want our singleton implement the interface so that we can test it with other concrete implementation)
Problems with singleton:
- often breaks the dependency inversion principle (depend on an interface/abstraction and not a concrete type)

## 5. ADAPTER
A construct which adapts an existing interface X to conform to the required interfaceY.
We use a special device to give us the interface we require from the interface we have.

* Determine the API you have and the API you need
* Create a component which aggregates (has a pointer to, ...) the adaptee
* Intermediate representations can pile up: use caching and other optimizations

## 6. BRIDGE
A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy).

It tries to avoid a compleity explosion.

- Prevents a Cartesian Product complexity explosion
- Example:
    * Common type ThreadScheduler
    * Can be preemptive or cooperative
    * Can run on Windows or LinuxS
    * End up with a 2x2 scenario: WindowsPTS, UnixPTS, WindowsCTS, UnixCTS

Separation of hierarchies.
A stronger form of encapsulation.

## 7. COMPOSITE
A mechanism for treating individual (scalar) objects and compositions of objects in a uniform manner.

Motivation:
- objects use other objects' fields/methods through embedding
- composition lets us make compund objects
    * e.g., a mathematical expression composed of simple expressions; or
    * a shape group made of several different shapes
- composite design pattern is used to treat both single (scalar) and composite objects uniformly
    (they would have the same interface)
    * i.e., Foo and []Foo have a common API

* some composed and singular objects need identical behaviors
* composite pattern allows us treat both singular and composed objects uniformly guarding them behind the interface
* Iteration supported with the Iterator design pattern 

## 8. DECORATOR
Facilitates the addition of behaviors to individual objects through embedding.

Motivation:
- Want to augment an object with additional functionality
- Do not want to rewrite or alter existing code (OCP)
- Want to keep new functionality separate (SRP)
- Need to be able to interact with existing structures
- Solution: embed the decorated object and provide additional functionality

* Decorator embeds the decorated object
* Adds utility fields and methods to augment the object's features
* Often used to emulate multiple inheritance (may require extra work)

## 9. FACADE
Provides a simple, easy to understand/use interface over a large and sophisticated body of code. 

Motivation:
- Balancing complexity and presentation/usability
- typical home:
    - many subsystems (electrical, sanitation)
    - comlpex internal structure (floor layers)
    - end users is not exposed to internals
- software:
    - many systems working to provide flexibility, but
    - API consumers want it to 'just work'

* Build a Facade to provide a simplified API over a set of components
* May wish to (optionally) expose internals through facade
* May allow users to 'escalate' to use more complex APIs if they need to