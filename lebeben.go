package main

import (
  "github.com/evanw/esbuild/pkg/api"
  "flag"
  "fmt"
  "os"
  "path/filepath"
  "time"
  "github.com/radovskyb/watcher"
)

// type Flags struct {
//   help bool
//   jsxFactory string
//   jsxFragment string
//   minify bool
//   nomodule bool
//   port int
//   public string
//   serve bool
//   watch string
//   verbose bool
// }

// [
//   {
//     long: "watch",
//     default: [],
//     type: "string[]",
//     description: "one or many directories to watch for changes",
//   },
//   {
//     long: "serve",
//     type: "boolean",
//     default: false,
//     description: "serve app locally",
//   },
//   {
//     long: "public",
//     short: "u",
//     type: "string",
//     default: "public",
//     description: "the directory to serve",
//   },
//   {
//     long: "port",
//     type: "number",
//     default: 5000,
//     description: "the port to serve",
//   },
//   {
//     long: "verbose",
//     type: "boolean",
//     default: false,
//     description: "enable more detailed logs",
//   },
//   {
//     long: "nomodule",
//     type: "boolean",
//     default: "false",
//     description: "generate the es2015 fallback bundle",
//   },
//   {
//     long: "minify",
//     type: "boolean",
//     default: false,
//     description: "minify all output",
//   },
//   {
//     long: "jsxFactory",
//     type: "string",
//     default: "h",
//     description: "jsx render function to use for nodes",
//   },
//   {
//     long: "jsxFragment",
//     type: "string",
//     default: "Fragment",
//     description: "jsx render function to use for fragments",
//   },
//   {
//     long: "help",
//     type: "boolean",
//     default: false,
//     description: "display this message",
//   },
// ];

func build(
  JsxFactory *string,
  JsxFragment *string,
  Minify *bool,
  Nomodule *bool,
  Public *string,
  Watch *string,
) (buildResult api.BuildResult) {
  Outfile := "module"
  Target := api.ES2017
  if *Nomodule {
    Target = api.ES2015
    Outfile = "nomodule"
  }
  buildResult = api.Build(api.BuildOptions{
    EntryPoints: []string{"src/app.jsx"},

    JSXFactory: *JsxFactory,
    JSXFragment: *JsxFragment,

    MinifyWhitespace:  *Minify,
    MinifyIdentifiers: *Minify,
    MinifySyntax:      *Minify,

    Target: Target,

    Outfile: filepath.Join(*Public, Outfile, "app.js"),
    Bundle:      true,
    Platform:    api.PlatformBrowser,
    Write:       true,
    Incremental: true,
  })
  fmt.Println("Built")
  return buildResult
}

func main() {
  Help := flag.Bool("help", false, "display this message")
  JsxFactory := flag.String("jsxFactory", "h", "jsx render function to use for nodes")
  JsxFragment := flag.String("jsxFragment", "Fragment", "jsx render function to use for fragments")
  Minify := flag.Bool("minify", false, "minify all output")
  Nomodule := flag.Bool("nomodule", false, "generate the es2015 fallback bundle")
  // Port := flag.Int("port", 5000, "the port to serve")
  Public := flag.String("public", "public", "the directory to serve")
  // Serve := flag.Bool("serve", false, "serve app locally")
  Watch := flag.String("watch", "", "one ~~or many~~ directories to watch for changes")
  // Verbose := flag.Bool("verbose", false, "enable more detailed logs")

  flag.Parse()

  if *Help {
    flag.PrintDefaults()
    os.Exit(1)
  }

  buildResult := build(
    JsxFactory,
    JsxFragment,
    Minify,
    Nomodule,
    Public,
    Watch,
  )

  if *Watch != "" {
    fmt.Println("Watching...")
    w := watcher.New()

    go func() {
      for {
        select {
        case event := <-w.Event:
          fmt.Println(event) // Print the event's info.
          // build(
          //   JsxFactory,
          //   JsxFragment,
          //   Minify,
          //   Nomodule,
          //   Public,
          //   Watch,
          // )
          buildResult.Rebuild()
        case err := <-w.Error:
          fmt.Println(err)
        case <-w.Closed:
          return
        }
      }
    }()

    if addErr := w.AddRecursive(*Watch); addErr != nil {
      fmt.Println("AddRecursive")
      fmt.Println(addErr)
    }

    if startErr :=  w.Start(time.Millisecond * 333); startErr != nil {
      fmt.Println("Start")
      fmt.Println(startErr)
    }
  }

  if len(buildResult.Errors) > 0 {
    fmt.Println("errors?")
    fmt.Println(buildResult.Errors)
  }
}
