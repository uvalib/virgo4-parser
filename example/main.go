package main

import (
	"log"

	"github.com/uvalib/virgo4-parser/v4parser"
)

/**
 * MAIN
 */
func main() {
	log.Printf("Testing out the validtaion behavior...")
	simple := "title: {bannanas}"
	validator := v4parser.Validator{}
	valid, errors := validator.Validate(simple)
	if valid == false {
		log.Printf("ERROR: [%s] is not valid, but is should be: %s", simple, errors)
	} else {
		log.Printf("SUCCESS: [%s] is valid", simple)
	}

	bad := "alligator: {bannana"
	valid, errors = validator.Validate(bad)
	if valid == false {
		log.Printf("SUCCESS: [%s] is not valid: %s", bad, errors)
	} else {
		log.Printf("ERROR: [%s] is valid, but it is not", bad)
	}

	test := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	valid, errors = validator.Validate(test)
	if valid == false {
		log.Printf("ERROR: [%s] is not valid, but should be: %s", test, errors)
	} else {
		log.Printf("SUCCESS: [%s] is valid", test)
	}

	log.Printf("Testing out SOLR conversion behavior...")
	solrOut := v4parser.ConvertToSolr(test)
	log.Printf("Result: %s", solrOut)
}
