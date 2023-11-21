package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (ch *CatalogHandler) DeleteItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	if err := ch.App.S.DeleteItem(itemID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete item"})
	}

	return c.JSON(fiber.Map{"message": "Item deleted successfully"})
}

/*func AddItemHandler(c *fiber.Ctx) error {
	var input NewItemInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	newItem := Item{
		Name:  input.Name,
		Price: input.Price,
	}
	if err := createItemInDB(newItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(newItem)
}

func UpdateItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	var updatedItem Item
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	err = updateItemInDB(itemObjectID, updatedItem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(updatedItem)
}
*/

/*
func GetItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	item, err := getItemFromDB(itemObjectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("404")
	}
	return c.JSON(item)
}
*/
