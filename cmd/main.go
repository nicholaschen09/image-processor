package main

import (
    "flag"
    "fmt"
    "image/jpeg"  // For encoding/decoding JPEG images
    "log"
    "os"        // Used for file operations
    "path/filepath"

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

    // Load the image
    img, err := utils.LoadImageFromPath(*inputPath)
    if err != nil {
        log.Fatalf("Error loading image: %v", err)
    }

    // Apply resize if specified
    if *width != 0 || *height != 0 {
        img = resize.Resize(img, *width, *height)
    }

    // Apply filter if specified
    if *filterType != "" {
        img = filters.ApplyFilter(img, *filterType)
    }

    // Create output directory if it doesn't exist
    outputDir := filepath.Dir(*outputPath)
    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
        if err := os.MkdirAll(outputDir, 0755); err != nil {
            log.Fatalf("Error creating output directory: %v", err)
        }
    }

    // Save the processed image
    outputFile, err := os.Create(*outputPath)
    if err != nil {
        log.Fatalf("Error creating output file: %v", err)
    }
    defer outputFile.Close()

    if err := jpeg.Encode(outputFile, img, nil); err != nil {
        log.Fatalf("Error encoding output image: %v", err)
    }

    fmt.Printf("Image successfully processed and saved to %s\n", *outputPath)
}