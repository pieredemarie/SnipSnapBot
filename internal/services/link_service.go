package services

import (
	"SnipSnapBot/internal/models"
	"SnipSnapBot/internal/repositories"
	"context"
	"time"
)

type LinkService struct {
	repo repositories.LinkRepository
}

type ILinkService interface {
	Save(ctx context.Context, userID int, url string, tags []string) error
	Edit(ctx context.Context, userID int, oldURL string, args []string) error
	GetByTag(ctx context.Context, userID int, tag string) ([]models.Link, error)
	Remove(ctx context.Context, userID int, url string) error
	List(ctx context.Context, userID int) ([]models.Link, error)
	GetRandom(ctx context.Context, userID int) (*models.Link, error)
}

func (s *LinkService) Save(ctx context.Context, userID int, url string, tags []string) error {
	if !s.isValidURL(url) {
		return ErrInvalidURL
	}

	if len(tags) == 0 {
		return ErrEmptyTags
	}

	newLink := models.Link{
		AuthorId: userID,
		URL:      url,
		Tags:     tags,
		Created:  time.Now().Unix(),
	}

	return s.repo.CreateLink(ctx, &newLink)
}

func (s *LinkService) List(ctx context.Context, userID int) ([]models.Link, error) {
	return s.repo.GetAllByUser(ctx, userID)
}

func (s *LinkService) GetRandom(ctx context.Context, userID int) (*models.Link, error) {
	return s.repo.GetRandom(ctx, userID)
}

func (s *LinkService) Remove(ctx context.Context, userID int, url string) error {
	return s.repo.DeleteLink(ctx, userID, url)
}

func (s *LinkService) Edit(ctx context.Context, userID int, oldURL string, args []string) error {
	if len(args) == 0 {
		return ErrNothingToEdit
	}

	var (
		newURL  *string
		newTags []string
	)

	// первый аргумент может быть URL
	first := args[0]

	if s.isValidURL(first) {
		newURL = &first
		newTags = args[1:]
	} else {
		newTags = args
	}

	if newURL == nil && len(newTags) == 0 {
		return ErrNothingToEdit
	}

	return s.repo.EditLink(ctx, userID, oldURL, newURL, &newTags)
}

func (s *LinkService) GetByTag(ctx context.Context, userID int, tag string) ([]models.Link, error) {
	return s.repo.GetByTag(ctx, userID, tag)
}
