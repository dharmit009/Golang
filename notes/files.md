# Files

* Linear access, not random access.
    ** Mechanical delay

## Basic Operation

| Operation Name | Functions               |
|----------------|-------------------------|
| Open           | Get handler for access. |
| Read           | Read bytes into []byte. |
| Write          | Write []byte into file. |
| Close          | Release handler.        |
| Seek           | Move read/write head.   |


## ioutil File Read

* io/ioutil package has basic functions.
* `dat` is `[]byte` filled with contents of entire file.
* No Explicit open/close are needed.
* Large files may cause problems as files are opened in memory.

**Reading a file:**

```golang
dat, err := ioutil.ReadFile("test.txt")

```

* **Writing a file:**

* Write `[]byte` to file.
* Creates a file.
* Unix-style permission bytes.

``` golang
dat = "Hello, World!"
err := ioutil.WriteFile("outline.txt", dat, 0777)

```




