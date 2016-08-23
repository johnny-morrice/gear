package cmd

import (
        "fmt"
        "os"
)

func die(err error) {
        fmt.Fprintf(os.Stderr, "fatal: %v", err)
        os.Exit(1)
}
