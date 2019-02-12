package main

import (
	"bitbucket.org/stulentsev/stackexchange"
	"fmt"
	"net/http"
	"os"
)

func main() {
	output := makePipeline(poller, filterer, onlyUnseen, markSeen)

	notifier(output)
}

func fetchLatestQuestions() ([]Question, error) {
	client := stackexchange.Client{
		Client:      http.DefaultClient,
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		Key:         "trkYWI2VTOWLi05pPpIulw((",
	}

	var questions []stackexchange.Question
	_, err := client.Do("/questions", &questions, &stackexchange.Params{
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
