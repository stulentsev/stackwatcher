package main

func tagsOverlap(tags1 []string, tags2 []string) bool {
  for _, a := range tags1 {
    for _, b := range tags2 {
      if a == b {
        return true
      }
    }
  }
  return false
}
