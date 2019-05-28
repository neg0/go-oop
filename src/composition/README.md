# Composition (Replacement for Inheritance)

Go does not support inheritance, however it does support composition. The generic definition of composition is "put together". 
Composition goes beyond the mechanics of type embedding. It’s a paradigm we can leverage to design better APIs and to build 
larger programs from smaller parts. It all starts from the declaration and implementation of types that have a single purpose. 
Programs that are designed with composition in mind have a better chance to grow and adapt to changing needs.

In conclusion, this example has tried to show how composition can be applied and the things to think about when writing Go code. How you declare your types and how they can work together is crucial.

* Declare the set of behaviors as discrete interface types first. Then think about how they can be composed into a larger set of behaviors.
* Only accept interface types for the behavior you are using in that function or method. This will help dictate the larger interface types that are required.
* Think about embedding as an inner and outer type relationship. Remember that through inner type promotion, everything that is declared in the inner type is promoted to the outer type. However, the inner type value exists in and of itself as is always accessible based on the rules for exporting.
* Type embedding is not sub-typing nor sub-classing. Concrete type values represent a single type and can’t be assigned based on any embedded relationship.
* The Go compiler can arrange interface conversions between related interface values. Interface conversion, at compile time, and it dismisses concrete types - it knows what to do merely based on the interface types themselves, not the implementing concrete values they could contain.