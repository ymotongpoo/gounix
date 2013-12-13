package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

const ProgramName = "whoami"

func main() {
	uid := os.Getuid()
	user, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		fmt.Printf("%s: cannot find name for user ID %d\n", ProgramName, uid)
	}
	fmt.Println(user.Username)
}
