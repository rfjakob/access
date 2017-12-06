package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("usage: access PATH\n")
		os.Exit(1)
	}
	path := flag.Arg(0)
	err := unix.Faccessat(unix.AT_FDCWD, path, unix.F_OK, unix.AT_SYMLINK_NOFOLLOW|1234)
	fmt.Println(err)
}
