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
    cd ~/go/src/github.com/mikaponics/mikapod-storage
    go run main.go
    ```

4. You now should see a message saying ``gRPC server is running`` then the application is running.

## Production
The following instructions are specific to getting setup for [Raspberry Pi](https://www.raspberrypi.org/).

### Deployment

1. (Optional) If already installed old golang with apt-get and you want to upgrade to the latest version. Run the following:

    ```
    sudo apt remove golang
    sudo apt-get autoremove
    source .profile
    ```

2. Install [Golang 1.11.8]():

    ```
    wget https://storage.googleapis.com/golang/go1.11.8.linux-armv6l.tar.gz
    sudo tar -C /usr/local -xzf go1.11.8.linux-armv6l.tar.gz
    export PATH=$PATH:/usr/local/go/bin # put into ~/.profile
    ```

3. Confirm we are using the correct version:

    ```
    go version
    ```

4. Install ``git``:

    ```
    sudo apt install git
    ```

5. Get our latest code.

    ```
    go get -u github.com/mikaponics/mikapod-storage
    ```

6. Install the depencies for this project.

    ```
    go get -u google.golang.org/grpc
    go get -u github.com/mattn/go-sqlite3
    ```

7. Go to our application directory.

    ```
    cd ~/go/src/github.com/mikaponics/mikapod-storage
    ```

8. (Optional) Confirm our application builds on the raspberry pi device. You now should see a message saying ``gRPC server is running`` then the application is running.

    ```
    go run main.go
    ```

9. Build for the ARM device and install it in our ``~/go/bin`` folder:

    ```
    go install
    ```

### Operation

1. While being logged in as ``pi`` run the following:

    ```
    sudo vi /etc/systemd/system/mikapod-storage.service
    ```

2. Copy and paste the following contents.

    ```
    [Unit]
    Description=Mikapod Storage Daemon
    After=multi-user.target

    [Service]
    Type=idle
    ExecStart=/home/pi/go/bin/mikapod-storage
    Restart=on-failure
    KillSignal=SIGTERM

    [Install]
    WantedBy=multi-user.target
    ```

3. We can now start the Gunicorn service we created and enable it so that it starts at boot:

    ```
    sudo systemctl start mikapod-storage
    sudo systemctl enable mikapod-storage
    ```

4. Confirm our service is running.

    ```
    sudo systemctl status mikapod-storage.service
    ```

5. If the service is working correctly you should see something like this at the bottom:

    ```
    raspberrypi systemd[1]: Started Mikapod Storage Daemon.
    ```

6. Congradulations, you have setup instrumentation micro-service! All other micro-services can now poll the latest data from the storage we have attached.

7. If you see any problems, run the following service to see what is wrong. More information can be found in [this article](https://unix.stackexchange.com/a/225407).

    ```
    sudo journalctl -u mikapod-storage
    ```

8. To reload the latest modifications to systemctl file.

    ```
    sudo systemctl daemon-reload
    ```

## License

This application is licensed under the **BSD** license. See [LICENSE](LICENSE) for more information.
