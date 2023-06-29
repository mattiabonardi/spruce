package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/db"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var customerCollection *mongo.Collection = db.OpenCollection(db.Client, "customers")
var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// create customer
func CreateCustomer(c *gin.Context) {

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(customer)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	customer.ID = primitive.NewObjectID()

	result, insertErr := customerCollection.InsertOne(ctx, customer)
	if insertErr != nil {
		msg := "customer item was not created"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

// get all custmers
func GetCustomers(c *gin.Context) {

	var customers []bson.M
	cursor, err := customerCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &customers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, customers)
}

// get customer by its id
func GetCustomerById(c *gin.Context) {

	customerID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(customerID)

	var customer bson.M

	if err := customerCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, customer)
}

// update the customer
func UpdateCustomer(c *gin.Context) {

	customerID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(customerID)

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(customer)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := customerCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		customer,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)
}

// delete a customer given the id
func DeleteCustomer(c *gin.Context) {

	customerID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(customerID)

	result, err := customerCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}
