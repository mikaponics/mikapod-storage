# Mikapod Storage
[![Go Report Card](https://goreportcard.com/badge/github.com/mikaponics/mikapod-storage)](https://goreportcard.com/report/github.com/mikaponics/mikapod-storage)

## Overview

The purpose of this application is to provide a local database tailored for storage of data for a **Mikapod-Soil** device and be accessible with remote procedure calls (RPC). The available operations are as follows:

* Create time-series datum

* List Time Series Data

* Delete time series data

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.12.7``

## Installation

1. Get our latest code.

    ```
    go get -u github.com/mikaponics/mikapod-storage
    ```

2. Install the depencies for this project.

    ```
    go get -u google.golang.org/grpc
    go get -u github.com/mattn/go-sqlite3
    ```

3. Run our application.

    ```
    cd github.com/mikaponics/mikapod-storage
    go run cmd/mikapod-storage/main.go
    ```

4. You now should see a message saying ``gRPC server is running`` then the application is running.


## License

This application is licensed under the **BSD** license. See [LICENSE](LICENSE) for more information.
