# Deallocating ...

Deallocating is a process of free up memory by removing or destroying
the variables, objects and other stuff which are no longer required.
This has to be done periodically else at some point you will run out
of memory which could make your entire application crash.

> This thing happens in C a lot which they refer to as **Memory Leak.**

Before we learn how to deallocate memory let us first talk about

## Where do we allocate memory?

There are two hunks of memory, namely:

* Stack: Primarily dedicated to **function calls**. You won't need to deallocate it.
* Heap: Persistent region of memory. You need to explicitly deallocate it.
