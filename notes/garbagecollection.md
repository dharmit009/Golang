# Garbage Collection:

**What is Garbage Collection?**

Garbage Collection is an automatic tool which is used to deal with
deallocation of memory. This is generally done in Interpreted
languages such as java & python. They does this by keeping track of
pointers and determines when the variable is in use and when it is
definitely not in use.

## Advantages:

* Easy for the programmer.
* Slow Needs Interpreted.

One Of the most unique feature of Garbage Collection in GO as it is a
complied type language which supports garbage collection by default.

The GO compiler is smart it determines whether it should allocate
space on stack vs on heap. The Garbage Collection is done in the
background.

## Disadvantages:

* A bit of a performance hit but nothing too pedantic.
