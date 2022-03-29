package helpers

import "log"

func Checkerror(err error) {
	if err != nil {
		log.Println(err)

	}
}
