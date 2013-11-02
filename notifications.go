package main

import (
  "fmt"
  "log"
  "math/rand"
  "os/exec"
  "strconv"
)

func notify(msg string, tags []string, url string) {
  notifyNC(msg, url)
  // notifyGrowl(msg, url)
}

func notifyNC(msg string, url string) {
  group := strconv.Itoa(rand.Int())
  cmd := exec.Command("terminal-notifier",
    "-title", "Greeting",
    "-group", group,
    "-message", fmt.Sprintf("Hello: %s", msg),
    "-sound", "Tink",
    "-open", url,
  )

  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
}

func notifyGrowl(msg string, url string) {
  // growlnotify -s -m "Hello world" --url "http://google.com" "my app"
  cmd := exec.Command("growlnotify",
    "--url", url,
    "-m", msg,
    "New Question",
  )

  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
}
