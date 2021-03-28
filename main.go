package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/lupinthe14th/sync/models"
	"github.com/lupinthe14th/sync/usecases"
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
			fmt.Printf("ID: %v Password: %v\n", u.ID, u.Password)
			if err := usecases.Echo(u.ID, u.Password); err != nil {
				log.Fatalf("echo failed: %v", err)
			}
			wg.Done()
		}(user)
	}
	wg.Wait()
}
