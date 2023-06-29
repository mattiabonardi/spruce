package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID           primitive.ObjectID `bson:"_id"`
	BusinessName *string            `json:"businessName"`
	Address      *string            `json:"address"`
	CAP          *string            `json:"cap"`
	Province     *string            `json:"province"`
	State        *string            `json:"state"`
	VAT          *string            `json:"vat"`
	Phone        *string            `json:"phone"`
}
