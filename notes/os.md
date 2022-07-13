# os package:

> ⚠️  **NOTE:**
> This is package is named `os` with a lowercase `o`.
| Operations | Functions                                    |
|------------|----------------------------------------------|
| os.open()  | Opens a file and return the file descriptor. |
| os.Read()  | Reads from a file into a []byte.             |
| os.Write   | Write to a file.                             |
| os.Close() | Closes a file.                               |

In os.Read() you can control the amount of data you want to read.

* Open and reading

```golang
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()

```

* Reads and fills `barr`.
* `Read()` returns # of bytes read.
* May be less than []byte length.

## os File Create / Write:

```golang
f, err := os.Create("outline.txt")
barr := [] byte { 1, 2, 3}
nb , err := f.Write(barr)
nb , err := f.WriteString("Hi")

```

* `WriteString()` writes a string.
* `Write()` writes a `[]byte`
    ** Any Unicode Sequence.

