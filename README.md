# Mikapod Storage
[![Go Report Card](https://goreportcard.com/badge/github.com/mikaponics/mikapod-storage)](https://goreportcard.com/report/github.com/mikaponics/mikapod-storage)

## Overview

The purpose of this application is to provide a local database tailored for storage of data for a **Mikapod-Soil** device and be accessible with remote procedure calls (RPC). The available operations are as follows:

* Create time-series datum

* List Time Series Data

* Delete time series data

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.15.6``

## Installation
1. Download the source code, build and install the application.

    ```
    GO111MODULE=on go get -u github.com/mikaponics/mikapod-storage
    ```

## Usage

Run our application.

    ```
    mikapod-storage serve
    ```

View the latest records

    ```
    mikapod-storage print --lines 100
    ```

3. You now should see a message saying ``Started storage service`` then the application is running.

## License

This application is licensed under the **BSD** license. See [LICENSE](LICENSE) for more information.
