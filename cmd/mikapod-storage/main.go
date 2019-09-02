package main // github.com/mikaponics/mikapod-soil/cmd/storage

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mikaponics/mikapod-storage/internal/app"
)


func main() {
	app := app.InitMikapodStorage()

    // DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
        <-sigs // Block execution until signal from terminal gets triggered here.
        app.StopMainRuntimeLoop()
    }()

    app.RunMainRuntimeLoop()
}
