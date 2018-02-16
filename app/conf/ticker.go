package conf

import "time"

func StartConfigManager() {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			LoadUserConfig()
		}
	}()
}
