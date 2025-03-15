package lib

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	err := BundleWebComponents("./components", "./static/index.js")
	if err != nil {
		fmt.Println(err.Error())
	}
}
