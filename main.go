package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/lupinthe14th/sync/models"
	"github.com/lupinthe14th/sync/usecases/doveadm"
)

func main() {
	var (
		users models.Users
		wg    sync.WaitGroup
	)
	if err := json.NewDecoder(os.Stdin).Decode(&users); err != nil {
		log.Fatalf("json decode error: %v", err)
	}
	for _, user := range users {
		wg.Add(1)
		go func(u *models.User) {
			if err := doveadm.Sync(u); err != nil {
				log.Fatalf("doveadm sync failed: %v", err)
			}
			wg.Done()
		}(user)
	}
	wg.Wait()
}
