package db

import (
	"fmt"
	"testing"
)

func Test_ConnectDB(t *testing.T) {
	client, err := connectDB(``)
	if err != nil {
		fmt.Println(err)
	}

	disconnectDB(client)
}
