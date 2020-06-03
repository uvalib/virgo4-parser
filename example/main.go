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
	valid, errors := v4parser.Validate(simple)
	if valid == false {
		log.Printf("ERROR: [%s] is not valid, but is should be: %s", simple, errors)
	} else {
		log.Printf("SUCCESS: [%s] is valid", simple)
	}

	bad := "alligator: {bannana"
	valid, errors = v4parser.Validate(bad)
	if valid == false {
		log.Printf("SUCCESS: [%s] is not valid: %s", bad, errors)
	} else {
		log.Printf("ERROR: [%s] is valid, but it is not", bad)
	}

	test := `( date:{1945/12/07 TO 1949} AND title : {"susan sontag" OR music title} AND keyword:{ Maunsell } ) OR author:{ liberty }`
	valid, errors = v4parser.Validate(test)
	if valid == false {
		log.Printf("ERROR: [%s] is not valid, but should be: %s", test, errors)
	} else {
		log.Printf("SUCCESS: [%s] is valid", test)
	}

	log.Printf("\n\nTesting out SOLR conversion behavior...")
	solrOut, err := v4parser.ConvertToSolr(simple)
	if err != nil {
		log.Printf("FAIL: %s", err.Error())
	} else {
		log.Printf("Result: %s", solrOut)
	}

	solrOut, err = v4parser.ConvertToSolr(test)
	if err != nil {
		log.Printf("FAIL: %s", err.Error())
	} else {
		log.Printf("Result: %s", solrOut)
	}

	solrOut, err = v4parser.ConvertToSolr(bad)
	if err != nil {
		log.Printf("SUCCESS: Detected error %s", err.Error())
	} else {
		log.Printf("FAIL: Parsed invalid string to %s", solrOut)
	}
	
	pub := `published: {New York City}`
	solrOut, err = v4parser.ConvertToSolr(pub)
	if err != nil {
		log.Printf("FAIL: %s", err.Error())
	} else {
		log.Printf("Result: %s", solrOut)
	}

	filter := `keyword: {"Organic chemistry"}  AND filter:{source_f:"Hathi Trust Digital Library"}`
	solrOut, err = v4parser.ConvertToSolr(filter)
	if err != nil {
		log.Printf("FAIL: %s", err.Error())
	} else {
		log.Printf("Result: %s", solrOut)
	}


}
