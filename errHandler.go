package main

import (
   "fmt"
   "os"
)

func processErr(err error){
   if err != nil {
      fmt.Fprintf(os.Stderr, "Error: %v\n", err)
      os.Exit(3)
   }
}
