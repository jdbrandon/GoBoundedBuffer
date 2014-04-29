package main

import (
   "math/rand"
)

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

func producer(prods int, iters int, buf chan<- rune, done chan<- bool){
   doneCount := make(chan bool)
   
   for i := 0; i < prods; i++ {
     go produce(iters, buf, doneCount)
   }
   for i := 0; i < prods; i++ {
      <-doneCount
   }

   close(doneCount)
   done <- true
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
