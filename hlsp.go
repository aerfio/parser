package parser

import "fmt"

// Parse receives either a YAML or JSON AsyncAPI document. It parses the document and checks if it's valid AsyncAPI. Skips specification extensions and schemas validation. If validation fails, the Parser/Validator should trigger an error. Produces a beautified version of the document in JSON Schema Draft 07
func Parse(document string) (string, err error) {

	return
}

func convertToJSON(doc string) (string, error) {
	if isJSON(doc) {
		// convert here JSON to YML
		return "", nil
	}

	if isYML(doc) {
		return doc, nil
	}

	return "", fmt.Errorf("Not supported Document format. Supported formats: JSON or YAML")
}

func isJSON(doc string) bool {
	return false
}

func isYML(doc string) bool {
	return false
}
