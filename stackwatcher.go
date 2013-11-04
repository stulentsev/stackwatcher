package main

import (
  "bitbucket.org/stulentsev/stackexchange"
  "fmt"
)


func main() {
  output := makePipeline(poller, filterer, onlyUnseen, markSeen)

  notifier(output)
}

func fetchLatestQuestions() ([]Question, error) {
  var questions []stackexchange.Question
  _, err := stackexchange.Do("/questions", &questions, &stackexchange.Params{
    Site:  stackexchange.StackOverflow,
    Sort:  stackexchange.SortCreationDate,
    Order: "desc",

    PageSize: 50,
  })

  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  var result []Question = make([]Question, len(questions))
  for i, question := range questions {
    result[i] = Question{
      ID:    question.ID,
      Title: question.Title,
      Tags:  question.Tags,
      Link:  question.Link,
    }
  }
  return result, nil
}
