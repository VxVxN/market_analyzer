package tools

import "github.com/VxVxN/log"

type closer interface {
	Close() error
}

func Close(file closer, message string) {
	if err := file.Close(); err != nil {
		log.Error.Printf("%s: %v", message, err)
	}
}

func ContainNumberInSlice(number int, slice []int) bool {
	for _, elem := range slice {
		if number == elem {
			return true
		}
	}
	return false
}
