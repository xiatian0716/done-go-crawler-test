package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	// ItemSaver里面就可以做事情
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %s", itemCount, item)
			itemCount++
		}
	}()
	return out
}
