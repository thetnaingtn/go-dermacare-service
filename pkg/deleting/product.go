package deleting

import (
	"github.com/thetnaingtn/go-dermacare-service/pkg/adding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	adding.Product `bson:",inline"`
}
