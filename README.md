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


## Test of Time
I left this "tool" as an absolute mess. But I did it on purpose (I swear) to see how much of a pain in the ass it would be to work on when I come back. As expected, it wasn't very pleasant. Since there were no tests I couldn't even think about adding features or fixing bugs. I first had to do a refactoring to make things more readable and easier do grasp while touching as little logic as possible. Then I implemented an end2end test so I could be sure things have not fallen apart after the next change will occur.

Now I was able to get rid of the unwanted deep directory reading that took place ([filepath.Walk](https://golang.org/pkg/path/filepath/#WalkFunc)) and change it to a shallow reading procedure ([ioutil.ReadDir](https://golang.org/pkg/io/ioutil/#ReadDir)). Since I was able to run the tests and see whether everything was still working as before, this change demanded close to no effort.