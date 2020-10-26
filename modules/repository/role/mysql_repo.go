package role

import (
	"database/sql"

	"arkan-jaya/business"
	core "arkan-jaya/core/role"
)

//MySQLRepository The implementation of role.Repository object
type MySQLRepository struct {
	db *sql.DB
}

//NewMySQLRepository Generate mongo DB role repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

//FindRoleByID Find user based on given ID. Its return nil if not found
func (repo *MySQLRepository) FindRoleByID(ID string) (*core.Role, error) {
	var role core.Role

	selectQuery := `SELECT id, name, is_admin, created_at, modified_at,
		FROM roles i
		WHERE i.id = ?`

	err := repo.db.
		QueryRow(selectQuery, ID).
		Scan(
			&role.ID, &role.Name, &role.PermissionID,
			&role.CreatedAt, &role.CreatedBy,
			&role.ModifiedAt, &role.ModifiedBy,
			&role.Version)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &role, nil
}

//GetAll Find all roles based on given tag. Its return empty array if not found
func (repo *MySQLRepository) GetAll() ([]core.Role, error) {
	//TODO: if feel have a performance issue in tag grouping, move the logic from db to here
	selectQuery := `SELECT id, name, is_admin, created_at, created_by, modified_at, modified_by, version
		FROM roles i`

	row, err := repo.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var roles []core.Role

	for row.Next() {
		var role core.Role

		err := row.Scan(
			&role.ID, &role.Name, &role.PermissionID,
			&role.CreatedAt, &role.CreatedBy,
			&role.ModifiedAt, &role.ModifiedBy,
			&role.Version)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err != nil {
		return nil, err
	}

	return roles, nil
}

//InsertRole Insert new role into database. Its return role id if success
func (repo *MySQLRepository) InsertRole(role core.Role) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	roleQuery := `INSERT INTO roles (
			id, 
			name, 
			is_admin, 
			created_at, 
			created_by, 
			modified_at, 
			modified_by,
			version
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	if err != nil {
		return err
	}

	_, err = tx.Exec(roleQuery,
		role.ID,
		role.Name,
		role.PermissionID,
		role.CreatedAt,
		role.CreatedBy,
		role.ModifiedAt,
		role.ModifiedBy,
		role.Version,
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

//UpdateRole Update existing role in database
func (repo *MySQLRepository) UpdateRole(role core.Role, currentVersion int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	roleUpdateQuery := `UPDATE roles
		SET
			name = ?,
			is_admin = ?,
			modified_at = ?,
			modified_by = ?,
			version = ?
		WHERE id = ? AND version = ?`

	res, err := tx.Exec(roleUpdateQuery,
		role.Name,
		role.PermissionID,
		role.ModifiedAt,
		role.ModifiedBy,
		role.Version,
		role.ID,
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

//DeleteRole Delete existing role in database
func (repo *MySQLRepository) DeleteRole(ID string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	roleDeleteQuery := `DELETE FROM role 
		WHERE id = ?`

	res, err := tx.Exec(roleDeleteQuery,
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
