package main

import "fmt"
func main() {

    famiFechaNac := make(map[string] int)

    famiFechaNac["Mama"] = 66
    famiFechaNac["Papa"] = 64

    fmt.Println(len(famiFechaNac))

    fmt.Println(famiFechaNac["Mama"])

    delete(famiFechaNac, "Mama")
    fmt.Println(len(famiFechaNac))
}
