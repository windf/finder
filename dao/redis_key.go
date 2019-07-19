package dao

import "fmt"

func RKRecord(id int64) string {
	return fmt.Sprintf("R:%d", id)
}
