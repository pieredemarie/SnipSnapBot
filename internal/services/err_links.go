package services

import "errors"

var (
	ErrNoLinks       = errors.New("no links found")
	ErrInvalidURL    = errors.New("invalid URL")
	ErrEmptyTags     = errors.New("no tags in message")
	ErrNothingToEdit = errors.New("nothing to edit")
)
