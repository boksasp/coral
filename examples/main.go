package main

import (
	"flag"
	"log"
	"time"

	"github.com/boksasp/coral"
)

var (
	key      = ""
	value    = ""
	input    = ""
	blankVal = false
	nullVal  = false
	print    = false
)

func init() {
	log.SetFlags(0)

	flag.StringVar(&key, "key", "", "Key to filter on. Format: node.subnode.subsubnode")
	flag.StringVar(&value, "value", "", "Key should match this value to be filtered to output.\nUse '-blankVal' if matching on empty string values.")
	flag.StringVar(&input, "json", "", "JSON input string")
	flag.BoolVar(&blankVal, "blankVal", false, "Set if the filter should match on blank values")
	flag.BoolVar(&nullVal, "nullVal", false, "Set if the filter should match on null values")
	flag.BoolVar(&print, "print", false, "Print json structure to output if match is found.")
	flag.Parse()

	if key == "" {
		log.Fatal("Must provide '-key' argument")
	}

	if input == "" {
		log.Fatal("Must provide JSON string with '-json'")
	}

	if value == "" && !blankVal {
		log.Fatal("Must provide '-blankVal=true' if '-value' is omitted.")
	}
}

func main() {
	start := time.Now()

	match, err := coral.Filter(input, key, value, nullVal)

	if err != nil {
		log.Fatal("Something failed during filtering:", err)
	}

	if match {
		log.Printf("🟢 match")
		if print {
			log.Println(input)
		}
	} else {
		log.Println("🔴 no match")
	}

	log.Printf("Time spent: %v", time.Since(start))
}
