package material

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Material struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Code        string             `bson:"code" json:"code"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Company     primitive.ObjectID `bson:"company" json:"company"`
	CreatedBy   primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedBy   primitive.ObjectID `bson:"updatedBy" json:"updatedBy"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
