## ABOUT
This script generates a `.go` file with strings of all declarations within a go package.

## MOTIVATION
idk maybe ill need this

## USAGE
```
go run main.go -input code/ -output ./stringified_decls.go
```
## INPUT
`code/main.go`
```
package main
  
import "fmt"
  
func main() {
    fmt.Println(add(1, 2))
}         
```
`code/helpers.go`
```
package main

func add(n1, n2 int) int {
    return n1 + n2
}
```

## OUTPUT
`./stringified_decls.go`
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

## FLAGS
|flag|usage|default|
|--|--|--|
|input|point to package to stringify|`./`|
|output|point to location of output file|`./stringified_decls.go`|
|package|define package name of output file|`main`|
|prefix|define prefix for declarations in output file| `""`|
|exclude|regular expression to match files to exclude| matches nothing|
|include|regular expression to match files to include| matches everything|
|only|string to equal the only file to include (eg. "main.go") | is ingored |

