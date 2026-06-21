# Batch Image Counter

A fast and efficient Go utility to count image files in directories, supporting recursive scanning with concurrent processing using worker goroutines.

## Features

- 📊 **Image File Counting** - Counts images by extension (.jpg, .jpeg, .png, .gif, .webp, .bmp)
- 🔄 **Recursive Scanning** - Optional recursive directory traversal
- ⚡ **Concurrent Processing** - Configurable worker goroutines for fast processing
- ⚙️ **YAML Configuration** - Simple configuration file support
- 📈 **Detailed Reports** - Displays file counts, folder statistics, and image totals

## Prerequisites

- Go 1.21 or higher
- YAML configuration file

## Installation

1. Clone the repository:
```bash
git clone https://github.com/pangdfg/Batch-Image-counter.git
cd Batch-Image-counter
```

2. Build the project:
```bash
go build -o batch-image-counter main.go
```

## Configuration

Create a `config/config.yml` file with the following structure:

```yaml
input_folder: "/path/to/your/images"
recursive: true              # true for recursive scanning, false for current directory only
workers: 8                   # number of worker goroutines
report_file: "report.csv"    # (optional) output report file
```

### Configuration Options

- **input_folder**: Path to the directory containing images
- **recursive**: Enable/disable recursive subdirectory scanning (boolean)
- **workers**: Number of concurrent worker goroutines (integer)
- **report_file**: Optional output file for the report

## Usage

### Basic Usage

```bash
./batch-image-counter
```

This will use the default configuration file at `config/config.yml`.

### Custom Configuration

```bash
./batch-image-counter -config /path/to/custom/config.yml
```

## Output

The tool prints a summary report including:

- Input folder path
- Recursive scanning status
- Number of workers used
- Image file counts by extension
- Total folder and file counts
- Total images found

Example output:
```
Input Folder : D:/Gallery
Recursive    : true
Workers      : 8

.jpg   : 150
.jpeg  : 45
.png   : 320
.gif   : 12
.webp  : 8
.bmp   : 5

Total Folders : 25
Total Files   : 540
Total Images  : 540
```

## Project Structure

```
.
├── main.go              # Main application entry point
├── go.mod              # Go module definition
├── go.sum              # Go dependencies lock file
├── README.md           # This file
├── config/
│   └── config.yml      # Configuration file (example)
├── models/
│   └── config.go       # Configuration data model
├── services/
│   └── loadconfig.go   # Configuration loading logic
└── response/
    └── res.go          # Response formatting and printing
```

## How It Works

1. **Configuration Loading**: Reads YAML configuration from `config/config.yml` or a custom path
2. **Directory Scanning**: Traverses the specified folder (recursively if enabled)
3. **Concurrent Processing**: Uses worker goroutines to process files concurrently
4. **Counting**: Counts files by extension and maintains thread-safe counters using mutex locks
5. **Reporting**: Prints comprehensive statistics about the scan results

## Performance

The tool uses:
- Worker goroutines for concurrent file processing
- Mutex locks for thread-safe counter updates
- Buffered channels for efficient file queue management
