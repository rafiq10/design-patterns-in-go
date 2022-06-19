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
- Some objects are simple and can be created in a single constructor call
- Other objects are more complicated to creae them at once
- Having a factory function with 10 arguments is not productive
- Instead, opt for piecewise (piece-by-piece) constructions
- Builder provides an API for constructing an object step-by-step

## 2.1 Creational Builder
- Fluent interfaces enable to chain calls together
- it returns the receiver object so you can chain it
- chaining the methods helps building the object

## 2.2 Builder Facets
- It splits the building of an object into multiple logical builder structs
- It allows, through the types, accessing the methods of builders that refer to the single fields
    of the principal struct
- Itallos chaining the methods to set the fields (in any order)