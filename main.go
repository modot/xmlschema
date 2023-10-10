package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if two file paths are passed as command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <xml_file_path> <xsd_file_path>")
		return
	}

	// Get the file paths from command line arguments
	xmlFilePath := os.Args[1]
	xsdFilePath := os.Args[2]

	// Load the XML file
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// Read the XML file into a byte slice
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	// Load the XSD schema
	xsdFile, err := os.Open(xsdFilePath)
	if err != nil {
		fmt.Println("Error opening XSD file:", err)
		return
	}
	defer xsdFile.Close()

	// Read the XSD schema into a byte slice
	xsdData, err := io.ReadAll(xsdFile)
	if err != nil {
		fmt.Println("Error reading XSD file:", err)
		return
	}

	// Parse the XSD schema
	schema := xmlschema.NewParser(bytes.NewReader(xsdData)).Parse()
	if err != nil {
		fmt.Println("Error parsing XSD schema:", err)
		return
	}

	// Parse the XML data
	var data interface{}
	err = xml.Unmarshal(xmlData, &data)
	if err != nil {
		fmt.Println("Error parsing XML data:", err)
		return
	}

	// Validate the XML data against the XSD schema
	err = schema.Validate(data)
	if err != nil {
		fmt.Println("XML data is not valid according to XSD schema:", err)
		return
	}

	fmt.Println("XML data is valid according to XSD schema.")
}
