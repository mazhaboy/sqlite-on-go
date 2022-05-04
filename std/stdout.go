package std

import (
	"fmt"
	"io"
	"os"
)

func stdout() {
	fmt.Println("Standard Output: os.Stdout")
	f := os.Stdout
	io.WriteString(f, "Hello this is a line 1 using io.Writing string\n")
	f.WriteString("This is line 2, using *File.WriteString\n")
	f.Write([]byte("This is line 3 using io.Writer interface\n"))
}
