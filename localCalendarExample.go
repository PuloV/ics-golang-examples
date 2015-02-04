package main

import (
	"fmt"
	"github.com/PuloV/ics-golang"
)

func main() {

	//  create new parser
	parser := ics.New()

	// get the input chan
	inputChan := parser.GetInputChan()

	//  send a local ics file
	inputChan <- "2eventsCal.ics"

	//  wait for the calendar to be parsed
	parser.Wait()

	// get all calendars in this parser
	cal, err := parser.GetCalendars()

	//  check for errors
	if err == nil {

		// print the first calendar like a string (calendar implements Stringer interface)
		fmt.Printf("%s \n", cal[0])

		// print the first calendar like a golang value
		fmt.Printf("%#v \n", cal[0])
	} else {
		// error
		fmt.Println(err)
	}

}
