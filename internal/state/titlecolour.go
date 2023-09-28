package state

import "time"

var TitleColourPeriod = 5

const (
	retropurplebright = "#b141f1"
	retropinkbright   = "#f92aad"
	retrogreenbright  = "#54e484"
	retroorangebright = "#ff7b00"
	retrobluebright   = "#58c7e0"
)

func TitleColourHandler(colorChan chan<- string, request <-chan struct{}) {
	colours := [5]string{
		retropurplebright,
		retropinkbright,
		retrogreenbright,
		retroorangebright,
		retrobluebright,
	}

	ticker := time.NewTicker(time.Duration(TitleColourPeriod) * time.Second)

	go func() {
		for {
			for _, colour := range colours {
				select {
				case <-ticker.C:
				case <-request:
					colorChan <- colour
				}
			}
		}
	}()
}
