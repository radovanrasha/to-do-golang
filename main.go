package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Reset  = "\033[0m"
)

func main() {
    fmt.Println(Green + "Welcome to to-do list app!" + Reset)
   	scanner := bufio.NewScanner(os.Stdin)
	var tasks []string
	
	for {
		fmt.Print(Yellow + "Enter your task (or type '--help' for help): " + Reset)
		if scanner.Scan() {
			userInput := scanner.Text()
			
			if strings.ToLower(userInput) == "--exit" {
				fmt.Println(Red + "Exiting program..." + Reset)
				os.Exit(0)
			} else if strings.ToLower(userInput) == "--help" {
				fmt.Println("type " + Yellow + "--list" + Reset + " - to list all tasks")
				fmt.Println("type " + Yellow + "--exit" + Reset + " - to exit app")
				fmt.Println("type " + Yellow + "--help" + Reset + " - to list all commands")
			}else if strings.ToLower(userInput) == "--list" {
				if len(tasks) == 0 {
					fmt.Println(Red + "No added tasks." + Reset)
				}
				for i := 0; i < len(tasks); i++ {
					fmt.Println(tasks[i])
				}
			}else{
				tasks = append(tasks, userInput)
				fmt.Printf("You added task: %s\n", userInput)
			}
		}
	}
}
