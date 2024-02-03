package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
}

type UsersRepository struct {
	c *mongo.Collection
}

func InitUserRepository(db *mongo.Database, collectionName string) IUserRepository {
	c := db.Collection(collectionName)
	return &UsersRepository{
		c,
	}
}

type IUserRepository interface {
	InsertOne(u User) (*primitive.ObjectID, error)
	FindByEmail(email string) (*User, error)
	FindById(id primitive.ObjectID) (*User, error)
}

func (r *UsersRepository) InsertOne(u User) (*primitive.ObjectID, error) {
	return nil, nil
}
func (r *UsersRepository) FindByEmail(email string) (*User, error) {
	return nil, nil
}

func (r *UsersRepository) FindById(id primitive.ObjectID) (*User, error) {
	return nil, nil
}
