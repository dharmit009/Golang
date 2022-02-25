# Reading Strings ...

Reading strings in Golang is a bit troublesome when compared to some
other programming languages particularly reading strings. 

**Example:**

```golang

import(
  "fmt", 
  "bufio", 
  "os",
  )

func main(){
  // To create a new reader we use the NewReader Function ...

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter a string: ")
  filename, _ := reader.ReadString('\n')
  fmt.Println("String: %s")

}
```
