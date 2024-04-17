package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

type user struct {
  name string
  age int
  Occupation string
}

func main() {
  fmt.Println("")
  file, err := os.Open("sample.txt")
  if err != nil {
    fmt.Println("Faylni ochishda xatolik yuz berdi:", err)
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var users []user
  for scanner.Scan() {
    line := scanner.Text()
    wordCount := []string{}

    words := strings.Fields(line)

    for _, i := range words {
      wordCount = append(wordCount, i)
    }
    if len(wordCount) > 0 {
      s := ""
      switch {

      case wordCount[0] == "Name:":
        for i:=1; i<len(wordCount); i++ {
          s += wordCount[i] + " "
        }
        users = append(users, user{s, 0, ""})
      case wordCount[0] == "Age:":
        users[len(users)-1].age, _ = strconv.Atoi(wordCount[1])
      case wordCount[0] == "Occupation:":
        for i:=1; i<len(wordCount); i++ {
          s += wordCount[i] + " "
        }
        users[len(users)-1].Occupation = s
      
      }
    }
    
  }
  fmt.Println(users)
}