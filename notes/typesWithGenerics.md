# Types with Generics:

I wanted to create a stack which could work with any type of input
so typically while creating a stack we need use:

`type Stack []int`

Now Let us see how we would define it using Generics.

`type Stack [T any] []T`

Let us try to understand the syntax in a bit more detail:

    * `type` is a keyword that is used to define a new type.
    * `Stack` is the name of the new type.
    * `[T any]` is a generic type parameter. This means that the Stack type can store any type of data, as long as it is a valid Golang type.
    * `[]T` is a slice of T objects. A slice is a Golang data structure that can store a dynamically-sized collection of elements.

In other words, the type Stack[T any] []T statement defines a new type called Stack that can store a slice of any type of data.
