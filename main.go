package main

import (
	"fmt"
	"os"
)

func main() {
	cmds := commands{make(map[string]func(cmd command) error)}
	cmds.register("monitor", handlerMonitor)

	input := os.Args
	fmt.Println(input)
	if len(input) < 2 {
		fmt.Fprint(os.Stderr, "Usage Error: stoplight <cmd name> [args...]")
		os.Exit(1)
	}

	cmd := command{
		Name: input[1],
		Args: input[2:],
	}

	err := cmds.run(command{Name: cmd.Name, Args: cmd.Args})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running: %s: %v", cmd.Name, err)
		os.Exit(1)
	}

}

