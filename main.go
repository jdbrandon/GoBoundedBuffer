package boundBuffer

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/* Global constants and variables */
const SIZE int = 200

/* Problem: Solve the bounded buffer problem using go.
   Solution: Create a buffer and producer and consumer
   routines that execute asynchronusly. Producers will
   place items in the buffer and consumers will remove
   the objects. Because no resource is infinite there
   is a size limit enforced on the buffer. It is
   important that a consumer does not read from an
   empty buffer and likewise a producer should not try
   to place an item into the buffer when it is full.
   Also because producers and consumers are acting
   asynchronusly it is important to enforce mutual
   exclusion on access to the buffer to avoid a race
   condition. The buffer is implemented as a channel.
   By using goroutines mutual exclusion is automatically
   enforced. Input is read from the command line.
   After operations on the channel are complete it is
   important to close them to return allocated
   resources to the system.
*/

func main() {

	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "Error: invalid input\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <nprods> <proditers> <ncons> <coniters>\n", os.Args[0])
		os.Exit(1)
	}

	nprods, err := strconv.ParseUint(os.Args[1], 0, 64)
	processErr(err)
	proditers, err := strconv.ParseUint(os.Args[2], 0, 64)
	processErr(err)
	ncons, err := strconv.ParseUint(os.Args[3], 0, 64)
	processErr(err)
	coniters, err := strconv.ParseUint(os.Args[4], 0, 64)
	processErr(err)

	if nprods*proditers != ncons*coniters {
		fmt.Fprintf(os.Stderr, "Error: nprods * proditers not equal to ncons * coniters\n")
		os.Exit(2)
	}

	rand.Seed(time.Now().Unix())
	buf := make(chan rune, SIZE)
	pdone := make(chan bool)
	cdone := make(chan bool)

	go consumer(int(ncons), int(coniters), buf, cdone)

	go producer(int(nprods), int(proditers), buf, pdone)

	select {
	case <-pdone:
		<-cdone
	case <-cdone:
		<-pdone
	}

	close(buf)
	close(pdone)
	close(cdone)

	os.Exit(0)
}
