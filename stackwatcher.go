package main

import (
  "bitbucket.org/stulentsev/stackexchange"
)

func main() {
  // scrapeQuestions()
}

func scrapeQuestions() error {
  var questions []stackexchange.Question
  _, err := stackexchange.Do("/questions", &questions, &stackexchange.Params{
    Site:     stackexchange.StackOverflow,
    Sort:     stackexchange.SortScore,
    Order:    "asc",
    PageSize: 3,
  })
  if err != nil {
    return err
  }

  for _, question := range questions {
    // fmt.Printf("%s (ID=%d)\n", question.Title, question.ID)
    notify(question.Title, question.Tags, question.Link)
  }
  return nil
}

