// cmd/quantumassets/main.go
package main

import (
"flag"
"log"
"os"

"quantumassets/internal/quantumassets"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumassets.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
