package role

import (
	"arkan-jaya/core/role"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoDBRepository The implementation of role.Repository object
type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	PermissionID    []string               `bson:"description"`
	CreatedAt  time.Time          `bson:"created_at"`
	CreatedBy  string             `bson:"created_by"`
	ModifiedAt time.Time          `bson:"modified_at"`
	ModifiedBy string             `bson:"modified_by"`
	Version    int                `bson:"version"`
}

func newCollection(role role.Role) (*collection, error) {
	objectID, err := primitive.ObjectIDFromHex(role.ID)

	if err != nil {
		return nil, err
	}

	return &collection{
		objectID,
		role.Name,
		role.PermissionID,
		role.CreatedAt,
		role.CreatedBy,
		role.ModifiedAt,
		role.ModifiedBy,
		role.Version,
	}, nil
}

func (col *collection) ToRole() role.Role {
	var role role.Role
	role.ID = col.ID.Hex()
	role.Name = col.Name
	role.PermissionID = col.PermissionID
	role.CreatedAt = col.CreatedAt
	role.CreatedBy = col.CreatedBy
	role.ModifiedAt = col.ModifiedAt
	role.ModifiedBy = col.ModifiedBy
	role.Version = col.Version

	return role
}

//NewMongoDBRepository Generate mongo DB role repository
func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("roles"),
	}
}

//FindRoleByID Find role based on given ID. Its return nil if not found
func (repo *MongoDBRepository) FindRoleByID(ID string) (*role.Role, error) {
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

	role := col.ToRole()
	return &role, nil
}

//GetAll Find all roles. Its return empty array if not found
func (repo *MongoDBRepository) GetAll() ([]role.Role, error) {
	filter := bson.M{}

	cursor, err := repo.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var roles []role.Role

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		roles = append(roles, col.ToRole())
	}

	return roles, nil
}

//InsertRole Insert new role into database. Its return role id if success
func (repo *MongoDBRepository) InsertRole(role role.Role) error {
	col, err := newCollection(role)

	if err != nil {
		return err
	}

	_, err = repo.col.InsertOne(context.TODO(), col)

	if err != nil {
		return err
	}

	return nil
}

//UpdateRole Update existing role in database
func (repo *MongoDBRepository) UpdateRole(role role.Role, currentVersion int) error {
	col, err := newCollection(role)

	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":     col.ID,
		"version": currentVersion,
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

//DeleteRole Delete existing role in database
func (repo *MongoDBRepository) DeleteRole(ID string) error {
	if len(ID) == 0 {
		return errors.New("Invalid ID")
	}

	filter := bson.M{
		"_id": ID,
	}

	_, err := repo.col.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
