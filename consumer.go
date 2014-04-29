package main

import (
   "fmt"
   "os"
)

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

func consumer(cons int, iters int, buf <-chan rune, done chan<- bool){
   doneCount := make(chan bool)

   for i := 0; i < cons; i++ {
      go consume(iters, buf, doneCount)
   }
   for i := 0; i < cons; i++ {
      <-doneCount
   }

   close(doneCount)
   done <- true
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
