package main

import "fmt"
import "strconv"
import "log"

func main() {
    var a string = "104"
    var b int = 35

    fmt.Printf("a data type:    %T\n", a)
    fmt.Printf("b data type: %T\n", b)

    c := strconv.Itoa(b)
    d, err := strconv.Atoi(a)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("-----------------")
    fmt.Printf("b data type:    %T\n", c)
    fmt.Printf("a data type: %T\n", d)

}

