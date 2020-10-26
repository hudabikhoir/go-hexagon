package permission

import (
	"database/sql"

	"arkan-jaya/business"
	core "arkan-jaya/core/permission"
)

//MySQLRepository The implementation of permission.Repository object
type MySQLRepository struct {
	db *sql.DB
}

//NewMySQLRepository Generate mongo DB permission repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

//FindPermissionByID Find user based on given ID. Its return nil if not found
func (repo *MySQLRepository) FindPermissionByID(ID string) (*core.Permission, error) {
	var permission core.Permission

	selectQuery := `SELECT id, resource, permission, created_at, modified_at,
		FROM permissions i
		WHERE i.id = ?`

	err := repo.db.
		QueryRow(selectQuery, ID).
		Scan(
			&permission.ID, &permission.Resource, &permission.Permission,
			&permission.CreatedAt, &permission.CreatedBy,
			&permission.ModifiedAt, &permission.ModifiedBy,
			&permission.Version)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &permission, nil
}

//GetAll Find all permissions based on given tag. Its return empty array if not found
func (repo *MySQLRepository) GetAll() ([]core.Permission, error) {
	//TODO: if feel have a performance issue in tag grouping, move the logic from db to here
	selectQuery := `SELECT id, resource, permission, created_at, created_by, modified_at, modified_by, version
		FROM permissions i`

	row, err := repo.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var permissions []core.Permission

	for row.Next() {
		var permission core.Permission

		err := row.Scan(
			&permission.ID, &permission.Resource, &permission.Permission,
			&permission.CreatedAt, &permission.CreatedBy,
			&permission.ModifiedAt, &permission.ModifiedBy,
			&permission.Version)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

//InsertPermission Insert new permission into database. Its return permission id if success
func (repo *MySQLRepository) InsertPermission(permission core.Permission) error {
	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	permissionQuery := `INSERT INTO permissions (
			id, 
			resource, 
			permission, 
			created_at, 
			created_by, 
			modified_at, 
			modified_by,
			version
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	if err != nil {
		return err
	}

	_, err = tx.Exec(permissionQuery,
		permission.ID,
		permission.Resource,
		permission.Permission,
		permission.CreatedAt,
		permission.CreatedBy,
		permission.ModifiedAt,
		permission.ModifiedBy,
		permission.Version,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

//UpdatePermission Update existing permission in database
func (repo *MySQLRepository) UpdatePermission(permission core.Permission, currentVersion int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	permissionUpdateQuery := `UPDATE permissions 
		SET
			resource = ?,
			permission = ?,
			modified_at = ?,
			modified_by = ?,
			version = ?
		WHERE id = ? AND version = ?`

	res, err := tx.Exec(permissionUpdateQuery,
		permission.Resource,
		permission.Permission,
		permission.ModifiedAt,
		permission.ModifiedBy,
		permission.Version,
		permission.ID,
		currentVersion,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		tx.Rollback()
		return err
	}

	if affected == 0 {
		tx.Rollback()
		return business.ErrZeroAffected
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

//DeletePermission Delete existing permission in database
func (repo *MySQLRepository) DeletePermission(ID string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	permissionDeleteQuery := `DELETE FROM permissions 
		WHERE id = ?`

	res, err := tx.Exec(permissionDeleteQuery,
		ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		tx.Rollback()
		return err
	}

	if affected == 0 {
		tx.Rollback()
		return business.ErrZeroAffected
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
