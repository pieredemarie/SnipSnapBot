package repositories

import (
	"SnipSnapBot/internal/models"
	"context"
	"errors"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrNoLinks = errors.New("no links found")
)

type LinkRepository interface {
	CreateLink(ctx context.Context, link *models.Link) error
	EditLink(ctx context.Context, userID int, oldURL string, newURL *string, tags *[]string) error
	GetAllByUser(ctx context.Context, userID int) ([]models.Link, error)
	GetByTag(ctx context.Context, userID int, tag string) ([]models.Link, error)
	GetRandom(ctx context.Context, userID int) (*models.Link, error)
	DeleteLink(ctx context.Context, userID int, url string) error
}

type MongoLinkRepo struct {
	collection *mongo.Collection
}

func NewMongoLinkRepo(collection *mongo.Collection) *MongoLinkRepo {
	return &MongoLinkRepo{
		collection: collection,
	}
}

func (m *MongoLinkRepo) CreateLink(ctx context.Context, link *models.Link) error {
	_, err := m.collection.InsertOne(ctx, link)
	return err
}

func (m *MongoLinkRepo) GetAllByUser(ctx context.Context, userID int) ([]models.Link, error) {
	var links []models.Link

	filter := bson.M{"author_id": userID}

	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &links); err != nil {
		return nil, err
	}

	if len(links) == 0 {
		return nil, ErrNoLinks
	}

	return links, err
}

func (m *MongoLinkRepo) GetByTag(ctx context.Context, userID int, tag string) ([]models.Link, error) {
	var links []models.Link

	filterID := bson.M{
		"author_id": userID,
		"tags":      tag,
	}

	cursor, err := m.collection.Find(ctx, filterID)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &links); err != nil {
		return nil, err
	}

	if len(links) == 0 {
		return nil, ErrNoLinks
	}

	return links, err
}

func (m *MongoLinkRepo) EditLink(ctx context.Context, userID int, oldURL string, newURL *string, tags *[]string) error {
	filter := bson.M{
		"author_id": userID,
		"url":       oldURL,
	}

	update := bson.M{}
	if newURL != nil {
		update["url"] = *newURL
	}
	if tags != nil {
		update["tags"] = *tags
	}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

func (m *MongoLinkRepo) DeleteLink(ctx context.Context, userID int, url string) error {
	filter := bson.M{
		"author_id": userID,
		"url":       url,
	}

	_, err := m.collection.DeleteOne(ctx, filter)
	return err
}

func (m *MongoLinkRepo) GetRandom(ctx context.Context, userID int) (*models.Link, error) {
	var links []models.Link

	filter := bson.M{
		"author_id": userID,
	}

	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &links); err != nil {
		return nil, err
	}

	if len(links) == 0 {
		return nil, ErrNoLinks
	}
	rand.Seed(time.Now().Unix())

	randomLink := rand.Intn(len(links))

	return &links[randomLink], nil
}
