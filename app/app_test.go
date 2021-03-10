package app_test

import (
	"context"
	"grader/db"
	"testing"
)

func TestAppNew(t *testing.T) {
	ctx := context.Background()
	dbClient, err := db.Connect(ctx)
	if err != nil {
		t.Error(err)
	}
	err = dbClient.Disconnect(ctx)
	if err != nil {
		t.Error(err)
	}
}
