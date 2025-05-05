package utils

import (
    "image"
    "os"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
)

// LoadImageFromPath loads an image from the specified path
// Different name to avoid conflict with existing LoadImage
func LoadImageFromPath(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    return img, err
}