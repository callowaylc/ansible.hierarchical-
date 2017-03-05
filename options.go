package hieransible

import(
  "os"
  "fmt"
  flags "github.com/jessevdk/go-flags"
)

// build inventory and build hosts file; maybe all in one..k
type Options struct {
  Inventory string `required:"true" default:"/etc/ansible/hosts" short:"i" long:"inventory-file" description:"specify inventory host path"`
  Install bool `long:"install" description:"installs hieransible inventory and host definitions"`
  Facts string `long:"facts" description:"determines facts for a given host"`
  Groups string `long:"groups" description:"determines groups for a given host"`
  Logs bool `long:"logs" description:"enables logging to stderr"`
}
var o Options

func InitOptions() {
  _, err := flags.Parse(&o)

  if flagsErr, ok := err.(*flags.Error); ok {
    switch flagsErr.Type {
    case flags.ErrHelp:
      os.Exit(0)
    case flags.ErrTag:
      panic(err)
    default:
      os.Exit(2)
    }

  } else if err != nil {
    panic(err)
  }
}

func GetOptions() Options {
  return o
}
