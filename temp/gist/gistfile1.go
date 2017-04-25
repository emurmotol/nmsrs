package main

import (
  "fmt"
  "os"
  "flag"
  "image"
  _ "image/gif"
  "image/jpeg"
  _ "image/png"
  "code.google.com/p/graphics-go/graphics"
)

func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) < 1 {
    fmt.Println("Input file is missing.");
    os.Exit(1);
  }
  if len(args) < 2 {
    fmt.Println("Output file is missing.");
    os.Exit(1);
  }
  fmt.Printf("opening %s\n", args[0])
  fSrc, err := os.Open(args[0])
  if err != nil {
      panic(err)
  }
  defer fSrc.Close()
  src, _, err := image.Decode(fSrc)
  if err != nil {
      panic(err)
  }
  dst := image.NewRGBA(image.Rect(0, 0, 80, 80))
  graphics.Thumbnail(dst, src)
  toimg, err := os.Create(args[1])
  if err != nil {
      panic(err)
  }
  defer toimg.Close()

  jpeg.Encode(toimg, dst, &jpeg.Options{jpeg.DefaultQuality})
}