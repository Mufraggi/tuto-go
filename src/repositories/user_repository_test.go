package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
	"time"
)

const (
	testMongoURIG     = "mongodb://localhost:27017"
	testDatabaseNameG = "testdb"
)

func setupTest(t *testing.T) (*mongo.Client, *mongo.Database, IUserRepository) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(testMongoURIG))
	assert.NoError(t, err)
	err = client.Ping(context.TODO(), nil)
	assert.NoError(t, err)
	database := client.Database(testDatabaseNameG)
	repo := InitUserRepository(database, "users")
	return client, database, repo
}

func deleteUserById(id primitive.ObjectID, database mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	database.Collection("users").DeleteOne(ctx, id)
}

func tearDownTest(client *mongo.Client, t *testing.T) {

	err := client.Disconnect(context.TODO())
	assert.NoError(t, err)
}

func TestInsertUser(t *testing.T) {
	client, db, repo := setupTest(t)
	defer tearDownTest(client, t)

	u := User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "email@gmail.com",
		Password:  "PassWord_hash",
	}

	objectId, err := repo.InsertOne(u)
	if err != nil {
		t.Fatalf("Error creating CompanyRepartition: %v", err)
	}
	if objectId == nil {
		t.Fatal("Returned ObjectID is nil")
	}
	deleteUserById(*objectId, *db)
}

func TestFindById(t *testing.T) {
	client, db, repo := setupTest(t)
	defer tearDownTest(client, t)
	u := User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "email@gmail.com",
		Password:  "PassWord_hash",
	}
	objectId, _ := repo.InsertOne(u)
	res, err := repo.FindById(*objectId)
	if err != nil {
		t.Fatalf("Error creating CompanyRepartition: %v", err)
	}
	if res.Password != u.Password {
		t.Fatalf("error pasword save")
	}
	if res.Email != u.Email {
		t.Fatalf("error pasword save")
	}
	deleteUserById(*objectId, *db)
}

func TestFindById_fail_notfound(t *testing.T) {
	client, _, repo := setupTest(t)
	defer tearDownTest(client, t)
	_, err := repo.FindById(primitive.NewObjectID())
	if err == nil {
		t.Fatalf("find by id dont return error")
	}
}

func TestFindByEmail(t *testing.T) {
	client, db, repo := setupTest(t)
	defer tearDownTest(client, t)
	u := User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     "email@gmail.com",
		Password:  "PassWord_hash",
	}
	objectId, _ := repo.InsertOne(u)
	resFindId, err := repo.FindById(*objectId)
	res, err := repo.FindByEmail(u.Email)
	if err != nil {
		t.Fatalf("Error creating CompanyRepartition: %v", err)
	}
	if !reflect.DeepEqual(resFindId, res) {
		t.Fatalf("error findByEmail")
	}
	deleteUserById(*objectId, *db)
}

func TestFindByEmail_fail_notfound(t *testing.T) {
	client, _, repo := setupTest(t)
	defer tearDownTest(client, t)
	_, err := repo.FindByEmail("qaa@gmail.com")
	if err == nil {
		t.Fatalf("find by id dont return error")
	}
}
