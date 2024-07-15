package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type completionObject struct {
	Title string		`json:"title"`
	Summary string	`json:"summary"`
	Comments string	`json:"comments"`
}

func (co completionObject)PrintObject(){
	fmt.Println("----------------------")
	fmt.Println("TICKET PROPOSAL")
	fmt.Println("")
	fmt.Println("TITLE:\t\t", co.Title)
	fmt.Println("SUMMARY:\t", co.Summary)
	fmt.Printf("")
	fmt.Println("----------------------")
}

func processCompletion(co completionObject) (completionObject, error){
	co.PrintObject()
	fmt.Print("\nDo you want to change title?(Y/N) ")
	reader := bufio.NewReader(os.Stdin)
	
			choice,_,_ := reader.ReadRune()
			if choice == 'Y' || choice == 'y'{
				fmt.Print("Give new title: ")
				reader := bufio.NewReader(os.Stdin)
				title,_ := reader.ReadString('\n')
				co.Title = strings.Replace(title, "\n", "", -1)
			}

			fmt.Print("\nDo you want to add something to summary? (Y/N) ")
			reader = bufio.NewReader(os.Stdin)
			choice,_,_ = reader.ReadRune()
			if choice == 'Y' || choice == 'y' {
				fmt.Print("Write text that will be added to summary: ")
				reader := bufio.NewReader(os.Stdin)
				summary,_ := reader.ReadString('\n')
				co.Summary = co.Summary + " " + summary
			}

	return co, nil
}