package main

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

func main(){

   if len(os.Args) != 5 {
      fmt.Fprintf(os.Stderr, "Error: invalid input\n")
      fmt.Fprintf(os.Stderr, "Usage: %s <nprods> <proditers> <ncons> <coniters>\n", os.Args[0])
      os.Exit(1)
   }

   nprods,err := strconv.ParseUint(os.Args[1], 0, 64)
   processErr(err)
   proditers,err := strconv.ParseUint(os.Args[2], 0, 64)
   processErr(err)
   ncons,err := strconv.ParseUint(os.Args[3], 0, 64)
   processErr(err)
   coniters,err := strconv.ParseUint(os.Args[4], 0, 64)
   processErr(err)

   if nprods * proditers != ncons * coniters {
      fmt.Fprintf(os.Stderr,"Error: nprods * proditers not equal to ncons * coniters\n")
      os.Exit(2)
   }

   rand.Seed(time.Now().Unix())
   buf := make(chan rune, SIZE)
   pdone := make(chan string)
   cdone := make(chan string)

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

/* Problem: Need to create producer threads that each produce a
   number of items to place in the buffer. 
   Given 4 parameters:
      prods -- number of producer threads to create
      iters -- number of items each producer should place in buffer
      buf   -- the buffer to place items into
      done  -- a chanel used to write to after production completes 
   Solution: Create prods goroutines, that procuce iters items
   wait for them to complete and then write a completion message 
   to the done channel.
*/

func producer(prods int, iters int, buf chan<- rune, done chan string){
   doneCount := make(chan bool)
   
   for i := 0; i < prods; i++ {
     go produce(iters, buf, doneCount)
   }
   for i := 0; i < prods; i++ {
      <-doneCount
   }

   close(doneCount)
   done <- "done producing"
}

/* Problem: Generate a random alphabetical rune and place it in
   the buffer a predetermined number of times.
   Given 3 parameters:
      iters -- number of items to place in buffer
      buf   -- the buffer
      done  -- a channel used to signal routine completion
   Solution: Generate a random character and write it to the
   buffer. Reapeat iters times. Write to done to signal completion.
*/

func produce(iters int, buf chan<- rune, done chan<- bool){
   for i :=0; i < iters; i++ {
      buf <- rune((rand.Float32() * 26) + 'A')
   }
   done <- true
}

/* Problem: Create consumer threads to remove items from the buffer
   Given 4 parameters:
      cons  -- the number of consumer threads to create
      iters -- the number of items each consumer should remove
      buf   -- the buffer to consume from
      done  -- channel used to signal completion
   Solution: Create cons goroutines, that consume iters items
   wait for them to complete and then write a completion message 
   to the done channel.
*/

func consumer(cons int, iters int, buf <-chan rune, done chan string){
   doneCount := make(chan bool)

   for i := 0; i < cons; i++ {
      go consume(iters, buf, doneCount)
   }
   for i := 0; i < cons; i++ {
      <-doneCount
   }

   close(doneCount)
   done <- "done consuming"
}

/* Problem: Consume a rune from a buffer a predetermined number 
   of times.
   Given 3 parameters:
      itr -- number of items to place in buffer
      buf   -- the buffer
      done  -- a channel used to signal routine completion
   Solution: Read a rune from the buffer and print it to stderr. 
   Reapeat itr times. Write to done to signal completion.
*/

func consume(itr int, buf <-chan rune, done chan<- bool){
   for i:= 0; i < itr; i++ {
      fmt.Fprintf(os.Stderr, "%c\n", <-buf)
   }
   done <- true
}

/* Problem: need to handle potential errors
   Solution: Check if the error exists. If it does print a
   relevant message to stderr and exit the program.
*/

func processErr(err error){
   if err != nil {
      fmt.Fprintf(os.Stderr, "Error: %v\n", err)
      os.Exit(3)
   }
}
