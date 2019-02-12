package main

import (
  "fmt"
  "log"
  "math/rand"
  "os"
  "os/exec"
  "strconv"
  "strings"
)

func notify(msg string, tags []string, url string) {
  if os.Getenv("ENABLE_TERMINAL_NOTIFIER") == "1" {
    notifyNC(msg, tags, url)
  }

  notifyTerminal(msg, tags)
}

func notifyTerminal(msg string, tags []string) {
  fmt.Printf("[%s] %s\n",
    strings.Join(tags, ", "),
    msg)
}

func notifyNC(msg string, tags []string, url string) {
  group := strconv.Itoa(rand.Int())
  cmd := exec.Command("terminal-notifier",
    "-title", strings.Join(tags, ", "),
    "-group", group,
    "-message", msg,
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
