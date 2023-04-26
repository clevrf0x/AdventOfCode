package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  // Print banner 
  fmt.Println("***************************")
  fmt.Println("* Advent Of Code          *")
  fmt.Println("* Day 1: Calorie Counting *")
  fmt.Println("***************************")
  
  // get file path from CLI Args
  if len(os.Args) < 2 {
    fmt.Println("You need to provide a input file path")
    fmt.Println("Syntax: ./main input")
    os.Exit(0)
  }
  filePath := os.Args[1]

  // Open input file
  file, err := os.Open(filePath)
  if err != nil {
    fmt.Println(err)
  }

  // Setup bufio scanner to read file lines one by one
  fileReader := bufio.NewScanner(file)
  fileReader.Split(bufio.ScanLines)

  totalCaloriesByElf := []int{0}

  // Read file lines and append it into fileLines array
  for fileReader.Scan() {
    convertedLineToInt, err := strconv.Atoi(fileReader.Text())
    // raises an error when program tries to convert empty line to int
    // in which case we can just skip the iteration and append a new element to array
    // else add convertedInt to last array item
    if err != nil {
      totalCaloriesByElf = append(totalCaloriesByElf, 0)
      continue
    }
    totalCaloriesByElf[len(totalCaloriesByElf) - 1] += convertedLineToInt 
  }
  
  // close opened file and free up memory
  file.Close()

  // print the highest number from totalCaloriesByElf list
  // set the first value as highest, then iterate through the list to find and replace
  // highest variable
  // highest := totalCaloriesByElf[0]
  // for i := 0; i < len(totalCaloriesByElf); i++ {
    // if highest < totalCaloriesByElf[i] {
      // highest = totalCaloriesByElf[i]
    // }
  // }

  // To solve gold star challenge, we need to find top three numbers fro totalCaloriesByElf list
  // and add that together, to simplify the code, i am going to comment out previous loop
  // sort the list in descending order and get the top three values
  // i will be writing simple bubble sorting algorithm
  arrayLength := len(totalCaloriesByElf)
  for i := 0; i < arrayLength; i++ {
    for j := 0; j < arrayLength - i - 1; j++ {
      if totalCaloriesByElf[j] < totalCaloriesByElf[j + 1] {
        // swaps variables
        totalCaloriesByElf[j], totalCaloriesByElf[j + 1] = totalCaloriesByElf[j + 1], totalCaloriesByElf[j]
      }
    }
  }
  fmt.Println("Highest Calory is: ", totalCaloriesByElf[0])
  topThree := totalCaloriesByElf[0] + totalCaloriesByElf[1] + totalCaloriesByElf[2]
  fmt.Println("Top Three Value Sum is: ", topThree)
}
