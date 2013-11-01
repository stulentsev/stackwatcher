package main

import (
  "bitbucket.org/stulentsev/stackexchange"
  "fmt"
  "log"
  "os/exec"
  "math/rand"
  "strconv"
)

func main() {
  SayHello() // stackwatcher.go:13: undefined: SayHello
  // but it's in types.go
}

// func scrapeQuestions() error {
//   var questions []stackexchange.Question
//   wrapper, err := stackexchange.Do("/questions", &questions, &stackexchange.Params{
//     Site:     stackexchange.StackOverflow,
//     Sort:     stackexchange.SortScore,
//     Order:    "asc",
//     PageSize: 3,
//   })
//   if err != nil {
//     return err
//   }
//   
//   for _, question := range questions {
//     fmt.Printf("%s (ID=%d)\n", question.Title, question.ID)
//     notify(question.Title, question.Link)
//   }
//   return nil
// }
// 
// func notify(msg string, url string) {
//   group := strconv.Itoa(rand.Int())
//   fmt.Println(group)
//   cmd := exec.Command("terminal-notifier",
//     "-title", "Greeting",
//     "-group", "a1",
//     "-message", fmt.Sprintf("Hello: %s", msg),
//     "-sound", "Ping",
//     "-open", url,
//   )
// 
//   err := cmd.Run()
//   if err != nil {
//     log.Fatal(err)
//   }
// }
