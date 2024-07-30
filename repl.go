package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type cliCommands struct{
    name string
    description string
    callback func() error
}

func startRepl(){
    scanner := bufio.NewScanner(os.Stdin)

    for{
        fmt.Print("Game >")
        scanner.Scan()
        input := scanner.Text()

        words := cleanInput(input)
        if len(words) <1{
            return
        }
        commandName := words[0]
        command,exists := get_commands()[commandName]

        if exists{
            var err error
            err = command.callback()
            if err !=nil{
                fmt.Println(err)
            }
        }else{
            fmt.Println("Unknown command, try help")
        }
    }
}

func cleanInput(s string)[]string{
    lowered := strings.ToLower(s)
    split := strings.Fields(lowered)
    return split
}

func get_commands()map[string]cliCommands{
    return map[string]cliCommands{
        "help":{
            name: "help",
            description: "Calls the help function",
            callback: commandHelp,
        },
        "exit":{
            name:"exit",
            description: "exits the program",
            callback: commandExit,
        },
        "start":{
            name:"start",
            description:"starts the quiz game",
            callback: commandStart,
        },
    }
}
