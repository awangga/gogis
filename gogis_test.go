/**
 * Author: Rolly Maulana Awangga
 * File: gogis_test.go
 */

package gogis

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := a + b

	if got := Add(a, b); got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestSubtract(t *testing.T) {
	a := 1
	b := 2
	expected := a - b

	if got := Subtract(a, b); got != expected {
		t.Errorf("Subtract(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestGetPublicIP(t *testing.T) {
	notfound := -1
	if got := strings.Index(GetPublicIP().Query, "."); got <= notfound {
		t.Errorf("Response Body : %d, didn't return json", got)
	}

}

func Test_UsingEnvvar(t *testing.T) {
	value := "mongodb+srv://gogis:gogis@cluster0.wghp85v.mongodb.net/?retryWrites=true&w=majority"
	os.Setenv("MONGO", value)
	if actual := os.Getenv("MONGO"); actual != value {
		t.Errorf("Env Value is : %v, didn't return actual value", actual)
	}

}

func TestGetLocation(t *testing.T) {
	var db MongoGeometry
	db.MongoString = "mongodb+srv://gogis:gogis@cluster0.wghp85v.mongodb.net/?retryWrites=true&w=majority"
	db.DBName = "location"
	db.CollectionName = "villages"
	db.LocationField = "border"
	desa := GetLocation(db, 107.57297533576039, -6.872079914985439)
	fmt.Sprintln(desa)
	if got := desa; got == nil {
		t.Errorf("Response Body : %v, didn't return json", got)
	}

}
