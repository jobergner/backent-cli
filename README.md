## ABOUT
This script generates a ".go" file with strings of all declarations within a go package.
## MOTIVATION
idk maybe ill need this

## INPUT
../code/main.go
```
package main
  
import "fmt"
  
func main() {
    fmt.Println(add(1, 2))
}         
```
../code/helpers.go
```
package main

func add(n1, n2 int) int {
    return n1 + n2
}
```

## USAGE
```
go run main.go -i ../code/ -o ./stringified_decls.go
```

## OUTPUT
./stringified_decls.go
```
package main

const main_import string = `import "fmt"`
const main_func string = `func main() {
    fmt.Println(add(1, 2))
}`
const add_func string = `func add(n1, n2 int) int {
    return n1 + n2
}
```