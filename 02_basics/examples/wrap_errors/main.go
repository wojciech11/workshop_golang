package main

import (
	"errors"
	"fmt"
)

var errAlertOn = errors.New("loud alert is on")

func disarmLock() error {
	// we tried
	// and failed :/
	return fmt.Errorf("cut wrong wire: %w", errAlertOn)
}

func breakSafe() error {
	err := disarmLock()
	return fmt.Errorf("failed to open the safe: %w", err)
}

func main() {
	err := breakSafe()

	if errors.Is(err, errAlertOn) {
		fmt.Printf("%+v", err)
	}
}
