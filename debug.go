package scraper

import (
	"fmt"
	"os"
)

// Enables debug json outputs
var Debug = false

func debugFileOutput(body []byte, format string, discriminator ...string) {
	if Debug {
		var template string
		if len(discriminator) > 0 {
			template = fmt.Sprintf(format, discriminator[0])
		} else {
			template = format
		}

		fmt.Printf("writing current api output to \"%s\"\n", template)
		os.WriteFile(template, body, 0777)
	}
}
