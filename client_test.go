// +build integration

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var key, secret, baseURL string

func init() {
	key = os.Getenv("KEY")
	secret = os.Getenv("SECRET")
	baseURL = "https://api.godaddy.com"
}

func TestInvalidUrl(t *testing.T) {
	_, err := NewClient("api.godaddy.com", key, secret)
	assert.NotNil(t, err)
}

func TestAuthFailure(t *testing.T) {
	client, err := NewClient(baseURL, "ABC", "123")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	_, err = client.GetDomainRecords("", "bogus.com")
	assert.NotNil(t, err)
}

func TestGetRecords(t *testing.T) {
	client, err := NewClient(baseURL, key, secret)
	assert.Nil(t, err)
	assert.NotNil(t, client)

	records, err := client.GetDomainRecords("", "n3integration.com")
	assert.Nil(t, err)
	assert.NotNil(t, records)

	for _, rec := range records {
		assert.NotNil(t, rec.Name)
		assert.NotNil(t, rec.Type)
		assert.NotNil(t, rec.Data)

		fmt.Println("REC -->", rec)
	}
}
