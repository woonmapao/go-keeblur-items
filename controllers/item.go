package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/woonmapao/go-keeblur-items/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mi MongoInstance

const dbName = "keeblurDB"

//offline test : replace with URI
const mongoURI = "mongodb://localhost:27017/" + dbName

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mi = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func GetAll(c *fiber.Ctx) error {

	query := bson.D{{}}

	cursor, err := mi.Db.Collection("items").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var items []models.Item = make([]models.Item, 0)

	if err := cursor.All(c.Context(), &items); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(items)
}

func GetByType(c *fiber.Ctx) error {

	itemType := c.Params("type")

	query := bson.D{primitive.E{Key: "item_type", Value: itemType}}

	cursor, err := mi.Db.Collection("items").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var items []models.Item = make([]models.Item, 0)

	if err := cursor.All(c.Context(), &items); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(items)
}

func GetByID(c *fiber.Ctx) error {

	idParam := c.Params("id")

	itemID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: itemID}}

	cursor, err := mi.Db.Collection("items").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var items []models.Item = make([]models.Item, 0)

	if err := cursor.All(c.Context(), &items); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(items)
}
