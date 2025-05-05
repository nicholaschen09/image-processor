package utils

import (
    "image"
    "image/jpeg"
    "image/png"
    "os"
)

func LoadImage(filePath string) (image.Image, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    return img, nil
}

func SaveImage(img image.Image, filePath string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    switch ext := filePath[len(filePath)-4:]; ext {
    case ".jpg", ".jpeg":
        return jpeg.Encode(file, img, nil)
    case ".png":
        return png.Encode(file, img)
    default:
        return nil
    }
}