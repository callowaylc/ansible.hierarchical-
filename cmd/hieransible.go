package main

import(
  "os"
  a "github.com/callowaylc/hieransible"
)

func init() {
  a.InitLogs()
  a.InitOptions()

  a.Logs("init hieransible", a.LogFields{
    "options": a.GetOptions(),
  })
}

func main() {
  options := a.GetOptions()

  os.Exit(0)
}
