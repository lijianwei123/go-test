package work

import (
	_ "fmt"
	_ "os"
)

type Search interface {
	search(content []string, findStr string) []string
}
