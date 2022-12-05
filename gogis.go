/**
 * Author: Rolly Maulana Awangga
 * File: gogis.go
 */

package gogis

import (
	"io"
	"net/http"
)

// Returns the sum of two numbers
func Add(a int, b int) int {
	return a + b
}

// Returns the difference between two numbers
func Subtract(a int, b int) int {
	return a - b
}

func GetPublicIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
