package main

import (
  "github.com/evanw/esbuild/pkg/api"
  "flag"
  "fmt"
  "net/http"
  "os"
  "path/filepath"
  "strings"
  "sync"
  "time"
  "github.com/radovskyb/watcher"
)

func build(
  EntryPoints []string,
  JsxFactory *string,
  JsxFragment *string,
  Minify *bool,
  Nomodule *bool,
  Public *string,
  Watch []string,
) (buildResult api.BuildResult) {
  startTime := time.Now()
  Outdir := "module"
  Target := api.ES2017
  if *Nomodule {
    Target = api.ES2015
    Outdir = "nomodule"
  }
  buildResult = api.Build(api.BuildOptions{
    EntryPoints: EntryPoints,

    JSXFactory: *JsxFactory,
    JSXFragment: *JsxFragment,

    MinifyWhitespace:  *Minify,
    MinifyIdentifiers: *Minify,
    MinifySyntax:      *Minify,

    Target: Target,

    Outdir: filepath.Join(*Public, Outdir),
    Bundle:      true,
    Platform:    api.PlatformBrowser,
    Write:       true,
    Incremental: true,
  })
  fmt.Println("⏱️  ", time.Since(startTime).Truncate(time.Millisecond))
  return buildResult
}

type MultiFlag []string
func (i *MultiFlag) String() string {
    return ""
}
func (i *MultiFlag) Set(value string) error {
    *i = append(*i, value)
    return nil
}
var Watch MultiFlag

func main() {
  Help := flag.Bool("help", false, "display this message")
  JsxFactory := flag.String("jsxFactory", "h", "jsx render function to use for nodes")
  JsxFragment := flag.String("jsxFragment", "Fragment", "jsx render function to use for fragments")
  Minify := flag.Bool("minify", false, "minify all output")
  Nomodule := flag.Bool("nomodule", false, "generate the es2015 fallback bundle")
  Port := flag.String("port", "5000", "the port to serve")
  Public := flag.String("public", "public", "the directory to serve")
  Serve := flag.Bool("serve", false, "serve app locally")
  flag.Var(&Watch, "watch", "one or many directories to watch for changes")
  // Verbose := flag.Bool("verbose", false, "enable more detailed logs")

  flag.Parse()

  if *Help {
    flag.PrintDefaults()
    os.Exit(1)
  }

  wg := new(sync.WaitGroup)

  buildResult := build(
    flag.Args(),
    JsxFactory,
    JsxFragment,
    Minify,
    Nomodule,
    Public,
    Watch,
  )

  if *Serve {
    wg.Add(1)
    go func() {
      http.Handle("/", http.FileServer(http.Dir(*Public)))
      http.ListenAndServe(":"+*Port, nil)
    }()
    fmt.Printf("Serving at http://localhost:%s\n", *Port)
  }

  if len(Watch) > 0 {
    fmt.Println("👀  Watching", strings.Join(Watch, ", "))
    w := watcher.New()

    go func() {
      for {
        select {
        case event := <-w.Event:
          startTime := time.Now()
          buildResult.Rebuild()
          fmt.Println("⏱️  ", time.Since(startTime).Truncate(time.Millisecond), event.Name())
        case err := <-w.Error:
          fmt.Println(err)
        case <-w.Closed:
          return
        }
      }
    }()

    for _, directory := range Watch {
      if addErr := w.AddRecursive(directory); addErr != nil {
        fmt.Println(addErr)
      }
    }

    if startErr :=  w.Start(time.Millisecond * 100); startErr != nil {
      fmt.Println(startErr)
    }
  }

  if len(buildResult.Errors) > 0 {
    fmt.Println("errors?")
    fmt.Println(buildResult.Errors)
  }

  wg.Wait()
}
