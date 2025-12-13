package tests

import (
	"SnipSnapBot/internal/services"
	"context"
	"testing"
)

func TestEdit_WithNewURLAndTags(t *testing.T) {
	repo := &MockRepo{}
	service := services.NewLinkService(repo)

	err := service.Edit(
		context.Background(),
		1,
		"old.com",
		[]string{"https://new.com", "go", "bot"},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.gotNewURL == nil || *repo.gotNewURL != "https://new.com" {
		t.Fatal("new URL was not set")
	}

	if len(*repo.gotTags) != 2 {
		t.Fatal("tags were not set correctly")
	}
}
