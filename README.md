# 1. SOLID

## 1.1 Single Responsability Pronciple
A type should have one primary responsabilty
* it should have one reason to change, the reason being related to its prime responsability
* an example of violating single responsability princile is a "God object"

## 1.2 Open-Closed Principle
Types should be:
 -  Open for extension
 -  Closed for modification

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
Stick to the go std library inerface style ("Reader","Writer","Closer", etc.)
Try to maintain the interfaces small and if they get bigger, just split them.

If we construct an interface with many methods and then we need a type that doesn't
have all the methods it will not pass as an argument to the function that expects the interface. 