package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrettyPrintJSON returns an object in neat JSON.
func PrettyPrintJSON(i interface{}) string {
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		fmt.Printf("could not marshal: %s\n", err.Error())
	}

	return string(b)
}

// WriteToFile writes an object to a file.
func WriteToFile(i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("could not marshal: %w", err)
	}

	err = os.WriteFile("output.json", b, 0o755)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
