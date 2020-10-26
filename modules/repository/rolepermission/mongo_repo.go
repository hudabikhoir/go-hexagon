package rolepermission

import (
	"arkan-jaya/core/rolepermission"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoDBRepository The implementation of role permission.Repository object
type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	RoleID       string    `bson:"role_id"`
	PermissionID string    `bson:"permission_id"`
	CreatedAt    time.Time `bson:"created_at"`
	CreatedBy    string    `bson:"created_by"`
	ModifiedAt   time.Time `bson:"modified_at"`
	ModifiedBy   string    `bson:"modified_by"`
	Version      int       `bson:"version"`
}

func newCollection(rolePermission rolepermission.RolePermission) (*collection, error) {
	return &collection{
		rolePermission.RoleID,
		rolePermission.PermissionID,
		rolePermission.CreatedAt,
		rolePermission.CreatedBy,
		rolePermission.ModifiedAt,
		rolePermission.ModifiedBy,
		rolePermission.Version,
	}, nil
}

func (col *collection) ToRolePermission() rolepermission.RolePermission {
	var rolePermission rolepermission.RolePermission
	rolePermission.RoleID = col.RoleID
	rolePermission.PermissionID = col.PermissionID
	rolePermission.CreatedAt = col.CreatedAt
	rolePermission.CreatedBy = col.CreatedBy
	rolePermission.ModifiedAt = col.ModifiedAt
	rolePermission.ModifiedBy = col.ModifiedBy
	rolePermission.Version = col.Version

	return rolePermission
}

//NewMongoDBRepository Generate mongo DB role permission repository
func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("role_permissions"),
	}
}

//FindRolePermissionByRoleID Find role permission based on given ID. Its return nil if not found
func (repo *MongoDBRepository) FindRolePermissionByRoleID(ID string) (*rolepermission.RolePermission, error) {
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

	rolePermission := col.ToRolePermission()
	return &rolePermission, nil
}

//GetAll Find all role permissions. Its return empty array if not found
func (repo *MongoDBRepository) GetAll() ([]rolepermission.RolePermission, error) {
	filter := bson.M{}

	cursor, err := repo.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var rolePermissions []rolepermission.RolePermission

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		rolePermissions = append(rolePermissions, col.ToRolePermission())
	}

	return rolePermissions, nil
}

//InsertRolePermission Insert new role permission into database. Its return role permission id if success
func (repo *MongoDBRepository) InsertRolePermission(rolePermission rolepermission.RolePermission) error {
	col, err := newCollection(rolePermission)

	if err != nil {
		return err
	}

	_, err = repo.col.InsertOne(context.TODO(), col)

	if err != nil {
		return err
	}

	return nil
}

//UpdateRolePermission Update existing role permission in database
func (repo *MongoDBRepository) UpdateRolePermission(rolePermission rolepermission.RolePermission, currentVersion int) error {
	col, err := newCollection(rolePermission)

	if err != nil {
		return err
	}

	filter := bson.M{
		"role_id":       col.RoleID,
		"permission_id": col.PermissionID,
		"version":       currentVersion,
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

//DeleteRolePermission Delete existing role permission in database
func (repo *MongoDBRepository) DeleteRolePermission(ID string) error {
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
