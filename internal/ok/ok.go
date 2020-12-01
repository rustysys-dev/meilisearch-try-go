package ok

import (
	"log"
)

func Ok(e error) {
	if e != nil {
		log.Fatal("ERROR: ", e)
	}
}
