package user

import (
	"arkan-jaya/core/user"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoDBRepository The implementation of user.Repository object
type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	Email      string             `bson:"email"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	RoleID     int  `bson:"role_id"`
	IsActive   []string `bson:"is_active"`
}

func newCollection(user user.User) (*collection, error) {
	objectID, err := primitive.ObjectIDFromHex(user.ID)

	if err != nil {
		return nil, err
	}

	return &collection{
		objectID,
		user.Name,
		user.Username,
		user.Password,
		user.Email,
		user.CreatedAt,
		user.ModifiedAt,
		user.RoleID,
		user.IsActive,
	}, nil
}

func (col *collection) ToUser() user.User {
	var user user.User
	user.ID = col.ID.Hex()
	user.Name = col.Name
	user.Username = col.Username
	user.Password = col.Password
	user.Email = col.Email
	user.CreatedAt = col.CreatedAt
	user.ModifiedAt = col.ModifiedAt
	user.RoleID = col.RoleID
	user.IsActive = col.IsActive

	return user
}

//NewMongoDBRepository Generate mongo DB user repository
func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("users"),
	}
}

//FindUserByID Find user based on given ID. Its return nil if not found
func (repo *MongoDBRepository) FindUserByID(ID string) (*user.User, error) {
	var col collection

	objectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		//if cannot be convert means that ID will be never found
		return nil, nil
	}

	filter := bson.M{
		"_id": objectID,
	}

	if err := repo.col.FindOne(context.TODO(), filter).Decode(&col); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	user := col.ToUser()
	return &user, nil
}

//FindAllUser Find all users based on given tag. Its return empty array if not found
func (repo *MongoDBRepository) FindAllUser() ([]user.User, error) {
	filter := bson.M{}

	cursor, err := repo.col.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var users []user.User

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		users = append(users, col.ToUser())
	}

	return users, nil
}

//InsertUser Insert new user into database. Its return user id if success
func (repo *MongoDBRepository) InsertUser(user user.User) error {
	col, err := newCollection(user)

	if err != nil {
		return err
	}

	_, err = repo.col.InsertOne(context.TODO(), col)

	if err != nil {
		return err
	}

	return nil
}

//UpdateUser Update existing user in database
func (repo *MongoDBRepository) UpdateUser(user user.User) error {
	col, err := newCollection(user)

	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": col.ID,
	}

	updated := bson.M{
		"$set": col,
	}

	_, err = repo.col.UpdateOne(context.TODO(), filter, updated)
	if err != nil {
		return err
	}

	return nil
}
