package exec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func execCommand() {
	fmt.Println("Hello\nworld\n!!!")
	cmd := exec.Command("echo", "I get it")
	cmd.Stdout = os.Stdout
	cmd.Run()
	time.Sleep(2 * time.Second)
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatalf("Command did not run successfully")
	}
	fmt.Printf("Output from ls: %s", out)
}
