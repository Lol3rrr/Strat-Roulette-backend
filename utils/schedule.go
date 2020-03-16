package utils

import "time"

// Schedule runs a function once and then every 'duration'
func Schedule(call func(), duration time.Duration) chan bool {
	stop := make(chan bool)

	ticker := time.NewTicker(duration)

	call()

	go func() {
		for {
			select {
			case <-ticker.C:
				call()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}