package tests

import (
	"SnipSnapBot/internal/models"
	"context"
)

type MockRepo struct {
	editCalled bool
	gotNewURL  *string
	gotTags    *[]string
}

func (m *MockRepo) EditLink(ctx context.Context, userID int, oldURL string, newURL *string, tags *[]string) error {
	m.editCalled = true
	m.gotNewURL = newURL
	m.gotTags = tags
	return nil
}

func (m *MockRepo) CreateLink(ctx context.Context, link *models.Link) error { return nil }
func (m *MockRepo) GetAllByUser(ctx context.Context, userID int) ([]models.Link, error) {
	return nil, nil
}
func (m *MockRepo) GetByTag(ctx context.Context, userID int, tag string) ([]models.Link, error) {
	return nil, nil
}
func (m *MockRepo) GetRandom(ctx context.Context, userID int) (*models.Link, error) {
	return nil, nil
}
func (m *MockRepo) DeleteLink(ctx context.Context, userID int, url string) error {
	return nil
}
