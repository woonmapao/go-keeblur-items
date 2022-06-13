package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/woonmapao/go-keeblur/models"
	"github.com/woonmapao/go-keeblur/storage"
	"gorm.io/gorm"
)

type Item struct {
	List_Item_Name             *string `json:"list_item_name"`
	List_Item_Type             *string `json:"list_item_type"`
	List_Item_Price            *int    `json:"list_item_price"`
	List_Item_Options_Colors   *string `json:"list_item_options_colors"`
	List_Item_Options_PCB      *string `json:"list_item_options_pcb"`
	List_Item_Options_Plate    *string `json:"list_item_options_plate"`
	List_Item_Options_Switches *string `json:"list_item_options_switches"`
	List_Item_Desc             *string `json:"list_item_desc"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetItemByType(context *fiber.Ctx) error {
	itemtype := context.Params("type")
	itemModel := &[]models.Items{}

	if itemtype == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "type cannot be empty"})
		return nil
	}

	fmt.Println("the type is ", itemtype)

	err := r.DB.Where("list_item_type = ?", itemtype).Find(itemModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get item(s)"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"massage": "item type fetched successfully",
			"data":    itemModel,
		})
	return nil
}

func (r *Repository) GetItemByID(context *fiber.Ctx) error {
	id := context.Params("id")
	itemModel := &models.Items{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	fmt.Println("the ID is ", id)

	err := r.DB.Where("id = ?", id).First(itemModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get item"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"massage": "item id fetched successfully",
			"data":    itemModel,
		})
	return nil
}

func (r *Repository) HelloWorld(context *fiber.Ctx) error {
	return context.SendString("Hello, World!")
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_type/:type", r.GetItemByType)
	api.Get("/get_id/:id", r.GetItemByID)
	app.Get("/", r.HelloWorld)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("Default port not specified")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":" + port)
}
