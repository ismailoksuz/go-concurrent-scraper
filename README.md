# Go Concurrent Web Scraper 🕷️

[![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🌟 Project Overview

This is a simple Command Line Interface (CLI) application built with **Go (Golang)** that demonstrates the power of **concurrent web scraping**.

The application fetches the `<title>` tags from a list of predefined URLs simultaneously. It is designed to showcase Go's high-performance concurrency model, which is ideal for I/O-bound tasks like networking.

### Key Concurrency Concepts:

| Concept | Go Implementation | Purpose |
| :--- | :--- | :--- |
| **Concurrency** | **Goroutines** (`go` keyword) | Launching multiple scraping tasks at once, independently and efficiently.  |
| **Communication** | **Channels** (`make(chan string)`) | Safely and synchronously collecting the results (page titles) back from all running tasks. |
| **Error Handling** | HTTP status codes and connection errors are handled to ensure fault tolerance. |

## 🛠️ How to Run

### Requirements

* Go (version 1.18 or higher)

### Execution Steps

1.  **Clone the Repository:**
    ```bash
    git clone [https://github.com/ismailoksuz/go-concurrent-scraper-.git](https://github.com/ismailoksuz/go-concurrent-scraper.git)
    cd go-concurrent-scraper-
    ```

2.  **Run the Application:**
    Since the project is a single file, you can run it directly using the `go run` command.

    ```bash
    go run main.go
    ```

### Example Output

The output sequence will be random because all requests start simultaneously and results are printed as soon as they are received.
