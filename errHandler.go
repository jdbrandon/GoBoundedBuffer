package boundBuffer

import (
	"fmt"
	"os"
)

/* Problem: need to handle potential errors
   Solution: Check if the error exists. If it does print a
   relevant message to stderr and exit the program.
*/

func processErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(3)
	}
}
