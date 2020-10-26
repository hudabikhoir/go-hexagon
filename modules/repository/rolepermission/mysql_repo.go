package rolepermission

import (
	"database/sql"

	"arkan-jaya/business"
	core "arkan-jaya/core/rolepermission"
)

//MySQLRepository The implementation of role permission.Repository object
type MySQLRepository struct {
	db *sql.DB
}

//NewMySQLRepository Generate mongo DB role permission repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

//FindRolePermissionByRoleID Find user based on given ID. Its return nil if not found
func (repo *MySQLRepository) FindRolePermissionByRoleID(roleID string) (*core.RolePermission, error) {
	var rolePermission core.RolePermission

	selectQuery := `SELECT role_id, permission_id, created_at, modified_at,
		FROM role_permissions i
		WHERE i.role_id = ?`

	err := repo.db.
		QueryRow(selectQuery, roleID).
		Scan(
			&rolePermission.RoleID, &rolePermission.PermissionID,
			&rolePermission.CreatedAt, &rolePermission.CreatedBy,
			&rolePermission.ModifiedAt, &rolePermission.ModifiedBy,
			&rolePermission.Version)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &rolePermission, nil
}

//GetAll Find all role_permissions based on given tag. Its return empty array if not found
func (repo *MySQLRepository) GetAll() ([]core.RolePermission, error) {
	//TODO: if feel have a performance issue in tag grouping, move the logic from db to here
	selectQuery := `SELECT role_id, permission_id, created_at, created_by, modified_at, modified_by, version
		FROM role_permissions i`

	row, err := repo.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var rolePermissions []core.RolePermission

	for row.Next() {
		var rolePermission core.RolePermission

		err := row.Scan(
			&rolePermission.RoleID, &rolePermission.PermissionID,
			&rolePermission.CreatedAt, &rolePermission.CreatedBy,
			&rolePermission.ModifiedAt, &rolePermission.ModifiedBy,
			&rolePermission.Version)

		if err != nil {
			return nil, err
		}

		rolePermissions = append(rolePermissions, rolePermission)
	}

	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

//InsertRolePermission Insert new rolePermission into database. Its return rolePermission id if success
func (repo *MySQLRepository) InsertRolePermission(rolePermission core.RolePermission) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	rolePermissionQuery := `INSERT INTO role_permissions (
			role_id, 
			permission_id, 
			created_at, 
			created_by, 
			modified_at, 
			modified_by,
			version
		) VALUES (?, ?, ?, ?, ?, ?, ?)`

	if err != nil {
		return err
	}

	_, err = tx.Exec(rolePermissionQuery,
		rolePermission.RoleID,
		rolePermission.PermissionID,
		rolePermission.CreatedAt,
		rolePermission.CreatedBy,
		rolePermission.ModifiedAt,
		rolePermission.ModifiedBy,
		rolePermission.Version,
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

//UpdateRolePermission Update existing role_permission in database
func (repo *MySQLRepository) UpdateRolePermission(rolePermission core.RolePermission, currentVersion int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	rolePermissionUpdateQuery := `UPDATE role_permissions
		SET
			role_id = ?,
			permission_id = ?,
			modified_at = ?,
			modified_by = ?,
			version = ?
		WHERE id = ? AND version = ?`

	res, err := tx.Exec(rolePermissionUpdateQuery,
		rolePermission.RoleID,
		rolePermission.PermissionID,
		rolePermission.ModifiedAt,
		rolePermission.ModifiedBy,
		rolePermission.Version,
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

//DeleteRolePermission Delete existing role_permission in database
func (repo *MySQLRepository) DeleteRolePermission(ID string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	rolePermissionDeleteQuery := `DELETE FROM role_permission 
		WHERE id = ?`

	res, err := tx.Exec(rolePermissionDeleteQuery,
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
