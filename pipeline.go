package main

import (
  "fmt"
  "log"
  "time"
)

type pipelineSegment func(in, out chan Question)

func makePipeline(segments ...pipelineSegment) (out chan Question) {
  var current_input chan Question
  var current_output chan Question

  for _, seg := range segments {
    current_output = make(chan Question)
    go seg(current_input, current_output)

    current_input = current_output
  }

  return current_output
}

func notifier(in chan Question) {
  for q := range in {
    notify(q.Title, q.Tags, q.Link)
  }
}

func poller(in, out chan Question) {
  for {
    select {
    case <-time.After(40 * time.Second):
      fmt.Println("Polling")
      questions, err := fetchLatestQuestions()

      if err != nil {
        log.Fatal(err)
      }

      for _, q := range questions {
        out <- q
      }
    }
  }
}

var cache = Cache{store: make(map[int]bool)}

func markSeen(in, out chan Question) {
  defer close(out)
  for q := range in {
    cache.MarkSeen(q.ID)
    out <- q
  }
}

func onlyUnseen(in, out chan Question) {
  defer close(out)
  for q := range in {
    if cache.HaveSeen(q.ID) == false {
      out <- q
    }
  }
}

func filterer(in, out chan Question) {
  defer close(out)
  for question := range in {
    my_tags := []string{"ruby", "ruby-on-rails", "mysql", "go", "mongodb"}

    if tagsOverlap(my_tags, question.Tags) {
      out <- question
    }
  }
}
