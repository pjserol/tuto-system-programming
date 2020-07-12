package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {
	fmt.Println("Current PID:", os.Getpid())
	fmt.Println("Current Parent PID:", os.Getppid())
	fmt.Println("User ID:", os.Getuid())
	fmt.Println("Group ID:", os.Getgid())

	groups, err := os.Getgroups()
	if err != nil {
		panic(err)
	}
	fmt.Println("Group IDs:", groups)

	uid := os.Getuid()
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %s (uid %d)\n", u.Username, uid)

	gid := os.Getgid()
	group, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Group: %s (gid %d)\n", group.Name, uid)
}
