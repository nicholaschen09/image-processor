# Image Processing Tool

This project is an image processing tool built in Go that allows users to apply various filters and resize images. It utilizes Go's image libraries and provides a simple command-line interface for processing images.

## Project Structure

```
image-processor
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── filters
│   │   ├── filters.go   # Common interfaces and types for image filters
│   │   ├── grayscale.go  # Grayscale filter implementation
│   │   └── sepia.go     # Sepia filter implementation
│   ├── resize
│   │   └── resize.go    # Image resizing implementation
│   └── utils
│       └── fileio.go    # Utility functions for file handling
├── assets
│   └── samples          # Sample images for testing
├── go.mod               # Module definition and dependencies
├── go.sum               # Module dependency checksums
└── README.md            # Project documentation
```

## Features

- **Apply Filters**: Supports applying grayscale and sepia filters to images.
- **Resize Images**: Allows resizing images to specified dimensions.
- **File Handling**: Load and save images from/to the filesystem.

## Usage

1. Clone the repository:
   ```
   git clone https://github.com/microsoft/vscode-remote-try-go.git
   cd image-processor
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go <image-path> <filter-type|resize-width> <resize-height>
   ```

   Replace `<image-path>` with the path to the image you want to process, `<filter-type>` with either `grayscale` or `sepia`, and `<resize-width>` and `<resize-height>` with the desired dimensions.

## Examples

- To apply a grayscale filter:
  ```
  go run cmd/main.go assets/samples/sample.jpg grayscale
  ```

- To resize an image:
  ```
  go run cmd/main.go assets/samples/sample.jpg resize 800 600
  ```

## Contributing

Feel free to submit issues or pull requests for improvements or additional features.