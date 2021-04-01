// Copyright © 2021 Hideo Suzuki
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/lupinthe14th/dovectl/models"
	"github.com/lupinthe14th/dovectl/pkg/doveadm"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Concurrency doveadm backup in Go",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("backup called")
		var (
			users models.Users
			wg    sync.WaitGroup
		)
		if err := json.NewDecoder(os.Stdin).Decode(&users); err != nil {
			return fmt.Errorf("json decode error: %v", err)
		}
		for _, user := range users {
			wg.Add(1)
			go func(u *models.User) {
				if err := doveadm.Backup(u); err != nil {
					fmt.Printf("doveadm backup failed: %v\n", err)
				}
				wg.Done()
			}(user)
		}
		wg.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
