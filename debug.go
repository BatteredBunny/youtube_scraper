package scraper

import (
	"fmt"
	"os"
)

// Enables debug json outputs
var Debug = false

// Configure the maximum length that the discriminator should display
const DebugDiscriminatorLength = 20

// Helper function for debug mode that outputs data to a file with specified name
func DebugFileOutput(body []byte, format string, discriminator ...string) {
	if Debug {
		var template string
		if len(discriminator) > 0 && len(discriminator[0]) <= DebugDiscriminatorLength {
			template = fmt.Sprintf(format, discriminator)
		} else if len(discriminator) > 0 {
			template = fmt.Sprintf(format, discriminator[0][:DebugDiscriminatorLength])
		} else {
			template = format
		}

		fmt.Printf("writing current api output to \"%s\"\n", template)
		os.WriteFile(template, body, 0777)
	}
}
