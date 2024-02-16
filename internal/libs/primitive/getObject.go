package primitive

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func  GetObject(id string) (primitive.ObjectID, error) {
	
  objID, err := primitive.ObjectIDFromHex(id)
  if err != nil {
      fmt.Println("Error:", err)
  }
	return objID, nil
}