package utils

import (
	"fmt"
	"io"
	"log"
)

func Close(closer io.Closer, source string) {
	err := closer.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("error closing: %s \n %e", source, err))
	}
}
