# JSON:

* JavaScript Object Notation
* RFC (7159)
* Commonly used data format.
* Used to capture structured format.
* Attribute - Value pair (which naturally corelates to struct or map.)


**Go Object Representation:**

``` go
p1 := Person(name: "Joe", addr: "Baker Street", phone: "123")

```

**JSON Data Representation:**

``` json
{"name": "Joe", "addr": "Baker Street", "phone": "123"}

```

## Json Properties:

* Json is all unicode.
* Human-readable.
* Fairly compact representation.
* Types can be combined recursively.
    ** i.e Arrays of structs, struct in struct.
* You can convert arbitratily Go objects into JSON objects.

## Json Marshalling:

**What do we mean by Marshalling?**

> The term Marshalling means converting objects into JSON representation.

``` golang
p1 := Person(name: "joe", addr: "adam street", phone:"123")
barr, err := json.Marshal(p1)

```

This upper piece of code will return JSON representation for the object
named `Person p1`. The `Marshal()` function will return JSON
representation as `[]byte` and `err`.

## JSON Unmarshalling:

JSON Unmarshalling is the process of converting JSON `[]byte`
representation into Go Object. **This is completely the opposite process
of JSON Marshalling.** We use the `Unmarshal()` function to achieve this
task.

``` golang
var p2 Person
err := json.Unmarshal(barr, &p2)

```

**Notice:** The first argument is barr i.e json representation and we need
to specify the object into which you want to store or unpack the data so
we need to provide a pointer the object p2.

Object must "fit" JSON []byte i.e. the Go object should have all the
attributes of JSON fields.

