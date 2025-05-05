package main

import (
    "flag"
    "fmt"
    "image"
    "log"
    "os"

    "image-processor/internal/filters"
    "image-processor/internal/resize"
    "image-processor/internal/utils"
)

func main() {
    inputPath := flag.String("input", "", "Path to the input image")
    outputPath := flag.String("output", "", "Path to save the processed image")
    filterType := flag.String("filter", "", "Type of filter to apply (grayscale, sepia)")
    width := flag.Int("width", 0, "Width to resize the image")
    height := flag.Int("height", 0, "Height to resize the image")

    flag.Parse()

    if *inputPath == "" || *outputPath == "" {
        log.Fatal("Input and output paths must be specified")
    }

    img, err := utils.LoadImage(*inputPath)
    if err != nil {
        log.Fatalf("Error loading image: %v", err)
    }

    if *filterType != "" {
        img = filters.ApplyFilter(img, *filterType)
    }

    if *width > 0 && *height > 0 {
        img = resize.Resize(img, *width, *height)
    }

    err = utils.SaveImage(img, *outputPath)
    if err != nil {
        log.Fatalf("Error saving image: %v", err)
    }

    fmt.Println("Image processed and saved successfully.")
}