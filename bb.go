package main

import (
   "fmt"
//   "math/rand"
   "os"
   "strconv"
)

/* Global constants */
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
   condition. The buffer is implemented as a circular
   array.  

   Note: I am still relatively new to go when this is
   being written so I am unaware of what structures
   are best and/or necessary for mutual exclusion. I 
   am planning on using goroutines for asynchronus 
   processes and Mutexes for enforcing mutual exclusion.
   I used gobyexample.com for many tutorials as I
   learned. This project is based on an assignment 
   I received in an operating systems course at Central
   Michigan University taught by Ishwar Rattan PhD. 
*/

func main(){

   if len(os.Args) != 5 {
      fmt.Printf("Error: invalid input\nUsage: %s <producer count> <producer iterations> <consumer count> <consumer iterations>\n", os.Args[0])
      os.Exit(1)
   }

   nprods,_ := strconv.ParseUint(os.Args[1], 0, 64)
   proditers,_ := strconv.ParseUint(os.Args[2], 0, 64)
   ncons,_ := strconv.ParseUint(os.Args[3], 0, 64)
   coniters,_ := strconv.ParseUint(os.Args[4], 0, 64)

   fmt.Printf("producers:\n\tc:%d\ti:%d\nconsumers:\n\tc:%d\ti:%d\n", nprods, proditers, ncons, coniters)

   var buf [SIZE]int
//   h := 0
//   t := 0

   /* create consumers */

   /* create producers */
   producer(&buf)
   fmt.Println(buf)
   /* wait for goroutines to complete */

   os.Exit(0)
}

func producer(buf *[SIZE]int){
   for i := 0; i < len(*buf); i++ {
      (*buf)[i] = i + 1
   }
}

func consumer(){

}
