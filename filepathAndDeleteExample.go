package main

import (
	"fmt"
	"github.com/PuloV/ics-golang"
)

func main() {

	//  create new parser
	parser := ics.New()

	// set the filepath for the ics files
	ics.FilePath = "tmp/new/"

	// we dont want to delete the temp files
	ics.DeleteTempFiles = false

	// get the input chan
	inputChan := parser.GetInputChan()

	// get the output chan => here we will expect all parsed events
	outputChan := parser.GetOutputChan()

	// send the calendar urls to be parsed
	inputChan <- "http://www.google.com/calendar/ical/bg.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"
	inputChan <- "http://www.google.com/calendar/ical/en.bulgarian%23holiday%40group.v.calendar.google.com/public/basic.ics"

	//  print events
	go func() {
		for event := range outputChan {
			fmt.Println(event.GetImportedID())
		}
	}()
	// wait to kill the main goroute
	parser.Wait()

}
