# BigO Calculations

Here is an example of how to calculate BigO Notation:

```javascript
function funChallenge(input){
    let a = 10;
    a = 50 +3;

    for(let i=0; i<input.length; i++){
        anotherfunction();
        let stranger = true;
        a++;
    }
    return a;
}

```

Let us try to break this stuff down and calculate it's BigO Notation. So the
first block of code results into **O(2)** as they are just assignment operator
which are constant nature in time.

```javascript
let a = 10 // O(1)
a = 50 + 3 // O(1)

```

The Second Block consist of a `for` loop, a function call, with some assignment &
incremental operation. This Block could be a little intimidating so let us break
this down further let us evaluate the code block within the `for` loop first.

```javascript
     for(let i=0; i<input.length; i++){
         anotherfunction();
         let stranger = true;
         a++;
     }

```

Now the cost of calling `anotherfunction()` from within `for` loop would be
**O(n)** as `anotherfunction()` would get called as many times as the `for` loop
keeps running. Therefore, the cost for `anotherfunction()` would be equal to
**O(N)**.

``` javascript
    anotherfunction();         // O(n)

```

The assignment and the incremental operations would also cost **O(n)** time due
to the same reasons mentioned earlier.

```javascript
    let stranger = true;  // O(n)
    a++;                  // O(n)

```

Lastly the `for` loop this would cost **O(n)** as it depends on the size of
inputs. So the final cost of the whole block would result into:

``` javascript
    for(let i=0; i<input.length; i++){     // O(n)
         anotherfunction();                // O(n)
         let stranger = true;              // O(n)
         a++;                              // O(n)
     }

```


Which is equivalent to **O(4n)**. And at the end there is one `return`
statement which result into **O(1)**.

``` javascript
return a; // O(1)

```

Now Let us see the final cost of this code:

```javascript
function funChallenge(input){

    let a = 10;
    a = 50 +3;                          // O(2)

    for(let i=0; i<input.length; i++){  // O(4n)
        anotherfunction();
        let stranger = true;
        a++;
    }

    return a;                           // O(1)
}

```

When you add all that together you get, **O(4n +3)**. Which later gets
simplified to just **O(n)**. You will learn that later.

> O(2) + O (4n) + O(1)
> = O (4n) + O (2) + O (1)
> = O (4n) + O (3)
> = O (4n + 3)


