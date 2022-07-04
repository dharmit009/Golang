# Concurrency:

**What is Concurrency?**

Concurrency is ability to execute multiples statements of code at the
same time. Concurrency is the management of multiple tasks at the same
time.

## Moore's Law:

Number of transistor doubles every 18months (this has quite slowed
down now).

This actually resulted into more transistor used to lead to higher
clock frequencies.

## Performance Limitation:

 Each machines have there performance limits which is highly depended on
 their CPU Hardware. The improvements in hardware are quite limited this
 is were concurrency comes in into picture.

## Parallelism:

* Number of core increases over time.
* Multiple tasks may be performed at the same time.

Difficulties with implementing parallelism:

* when to start / stop?
* What if one tasks needs data from another task?
* Do tasks conflict in memory?

## Concurrent Programming:

Concurrent programming is the management of multiple tasks at the same
time. It is also known as parallel programming.

1. Management of task execution.
1. Communication between tasks.
1. Synchronization between tasks.

## Concurrency in GO:

* Go includes concurrency primitives.
* `Goroutines` represent concurrent tasks.
* `Channels` are used to communicate between tasks.
* `Select` enables task synchronization.
* Concurrency primitives are efficient and easy to use.


