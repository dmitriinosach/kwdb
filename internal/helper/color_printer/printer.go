package cprntr

import "fmt"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

func PrintRed(m string) {
	fmt.Println(Red + m + Reset)
}

func PrintGreen(m string) {
	fmt.Println(Green + m + Reset)
}

func PrintYellow(m string) {
	fmt.Println(Yellow + m + Reset)
}
