package state

import (
	"sync"
	"time"
)

const (
	CarouselTick = 5
)

const (
	PhotoOne   = "photo_one"
	PhotoTwo   = "photo_two"
	PhotoThree = "photo_three"
)

type CarouselState struct {
	mu     sync.RWMutex
	Margin map[string]int
	Photo  map[string]string
}

func initState() CarouselState {
	return CarouselState{
		mu: sync.RWMutex{},
		Margin: map[string]int{
			PhotoOne:   3,
			PhotoTwo:   0,
			PhotoThree: 5,
		},
		Photo: map[string]string{
			PhotoOne:   "bulma-logo.png",
			PhotoTwo:   "go-logo.png",
			PhotoThree: "htmx-logo.png",
		},
	}
}

func Carouselhandler(stateChan chan<- CarouselState, request <-chan struct{}) {
	state := initState()
	ticker := time.NewTicker(CarouselTick * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				permutateCarousel(&state)
			case <-request:
				stateChan <- state
			}
		}
	}()
}

func permutateCarousel(state *CarouselState) {
	state.mu.Lock()
	defer state.mu.Unlock()

	// Permutate the margin
	photoOneMargin := state.Margin[PhotoOne]
	photoTwoMargin := state.Margin[PhotoTwo]
	photoThreeMargin := state.Margin[PhotoThree]

	state.Margin[PhotoOne] = photoThreeMargin
	state.Margin[PhotoTwo] = photoOneMargin
	state.Margin[PhotoThree] = photoTwoMargin

	// Permutate the photo
	photoOne := state.Photo[PhotoOne]
	photoTwo := state.Photo[PhotoTwo]
	photoThree := state.Photo[PhotoThree]

	state.Photo[PhotoOne] = photoThree
	state.Photo[PhotoTwo] = photoOne
	state.Photo[PhotoThree] = photoTwo
}
