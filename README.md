# Log Analyzer

This repository is written in Go. Please follow the below steps to clone and run it:

## Prerequisites
- Install [Go](https://golang.org/doc/install).

## After Cloning the Repository, Run it as follows:

1. **Install Dependencies**
    ```bash
    go mod tidy
    ```

2. **Run the Application in Debug mode**
    ```bash
    go run . --debug
    ```

3. **Run the Application in Live Server mode**
    ```bash
    go run . --endpoint <ENDPOINT> --days <DAYS>
    ```
