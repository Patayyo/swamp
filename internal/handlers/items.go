package handlers

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

func DeleteItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	collection := db.Database("proba").Collection("Items")
	filter := bson.M{"_id": itemObjectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

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
