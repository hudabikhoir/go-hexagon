package permission

import (
	"arkan-jaya/core/permission"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoDBRepository The implementation of permission.Repository object
type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	ID         primitive.ObjectID `bson:"_id"`
	Resource   string             `bson:"resource"`
	Permission string             `bson:"permission"`
	CreatedAt  time.Time          `bson:"created_at"`
	CreatedBy  string             `bson:"created_by"`
	ModifiedAt time.Time          `bson:"modified_at"`
	ModifiedBy string             `bson:"modified_by"`
	Version    int                `bson:"version"`
}

func newCollection(permission permission.Permission) (*collection, error) {
	objectID, err := primitive.ObjectIDFromHex(permission.ID)

	if err != nil {
		return nil, err
	}

	return &collection{
		objectID,
		permission.Resource,
		permission.Permission,
		permission.CreatedAt,
		permission.CreatedBy,
		permission.ModifiedAt,
		permission.ModifiedBy,
		permission.Version,
	}, nil
}

func (col *collection) ToPermission() permission.Permission {
	var permission permission.Permission
	permission.ID = col.ID.Hex()
	permission.Resource = col.Resource
	permission.Permission = col.Permission
	permission.CreatedAt = col.CreatedAt
	permission.CreatedBy = col.CreatedBy
	permission.ModifiedAt = col.ModifiedAt
	permission.ModifiedBy = col.ModifiedBy
	permission.Version = col.Version

	return permission
}

//NewMongoDBRepository Generate mongo DB permission repository
func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("permissions"),
	}
}

//FindPermissionByID Find permission based on given ID. Its return nil if not found
func (repo *MongoDBRepository) FindPermissionByID(ID string) (*permission.Permission, error) {
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

	permission := col.ToPermission()
	return &permission, nil
}

//GetAll Find all permissions. Its return empty array if not found
func (repo *MongoDBRepository) GetAll() ([]permission.Permission, error) {
	filter := bson.M{}

	cursor, err := repo.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var permissions []permission.Permission

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		permissions = append(permissions, col.ToPermission())
	}

	return permissions, nil
}

//InsertPermission Insert new permission into database. Its return permission id if success
func (repo *MongoDBRepository) InsertPermission(permission permission.Permission) error {
	col, err := newCollection(permission)

	if err != nil {
		return err
	}

	_, err = repo.col.InsertOne(context.TODO(), col)

	if err != nil {
		return err
	}

	return nil
}

//UpdatePermission Update existing permission in database
func (repo *MongoDBRepository) UpdatePermission(permission permission.Permission, currentVersion int) error {
	col, err := newCollection(permission)

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

//DeletePermission Delete existing permission in database
func (repo *MongoDBRepository) DeletePermission(ID string) error {
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
