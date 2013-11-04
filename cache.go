package main

import (
  "sync"
)

type Cache struct {
  store map[int]bool
  sync.RWMutex
}

func (c *Cache) HaveSeen(qid int) bool {
  c.RLock()
  defer c.RUnlock()
  
  _, ok := c.store[qid]
  return ok
}

func (c *Cache) MarkSeen(qid int) {
  c.Lock()
  defer c.Unlock()
  
  c.store[qid] = true
}