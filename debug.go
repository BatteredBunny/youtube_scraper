package scraper

import (
	"fmt"
	"os"
)

// Enables debug json outputs
var Debug = false

const debugDiscriminatorLength = 20

func debugFileOutput(body []byte, format string, discriminator ...string) {
	if Debug {
		var template string
		if len(discriminator) > 0 && len(discriminator[0]) <= debugDiscriminatorLength {
			template = fmt.Sprintf(format, discriminator)
		} else if len(discriminator) > 0 {
			template = fmt.Sprintf(format, discriminator[0][:debugDiscriminatorLength])
		} else {
			template = format
		}

		fmt.Printf("writing current api output to \"%s\"\n", template)
		os.WriteFile(template, body, 0777)
	}
}
