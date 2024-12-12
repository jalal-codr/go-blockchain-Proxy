package jobs

import (
	"fmt"
	"log"
	"time"
)

func RunJobs() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			body, err := MakeGetRequest()
			if err != nil {
				log.Println("Error:", err)
				continue
			}
			// Print the response body
			fmt.Println("Response Body:", body)
		}
	}
}
