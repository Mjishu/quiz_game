package main

import (
    "fmt"
    "strings"
)

func commandHelp(){
    fmt.Println()
    fmt.Println("Help function")
    fmt.Println(strings.Repeat("-",50))
    fmt.Println()
}

func commandExit(){
    fmt.Println("exits the program")
}

func commandStart(){
    fmt.Println("starts the program")
}
