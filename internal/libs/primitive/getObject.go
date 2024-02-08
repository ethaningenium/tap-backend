package primitive

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func  GetObject(id string) (primitive.ObjectID, error) {
	idx := strings.Index(id, "(")
  if idx == -1 {
      fmt.Println("Invalid string format")
      return primitive.NilObjectID, nil
  }
  idStr := id[idx+2 : len(id)-2]
  // Преобразование строки в ObjectID
  objID, err := primitive.ObjectIDFromHex(idStr)
  if err != nil {
      fmt.Println("Error:", err)    }
  // Вывод ObjectID
  fmt.Println("ObjectID:", objID)
	return objID, nil
}