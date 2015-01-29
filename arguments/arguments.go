package arguments

import (
	"os"
	"strings"
)

func Get(idx int) (bool, string) {
	args := os.Args[1:]

	if idx < len(args) {
		return true, args[idx]
	}

	return false, ""
}

func Flag(name string) bool {
	return Provided("-" + name)
}

func Provided(name string) bool {
	args := os.Args[1:]

	for _, arg := range args {
		if arg == name {
			return true
		}
	}

	return false
}

func Value(name string) (bool, string) {
	args := os.Args[1:]
	name = "--" + name

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, name+"=") {
			return true, strings.Join(strings.Split(arg, "=")[1:], "=")
		} else if arg == name && i+1 < len(args) {
			i := i + 1
			return true, args[i]
		}
	}

	return false, ""
}
