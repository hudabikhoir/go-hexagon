package user

import (
	"database/sql"

	"arkan-jaya/business"
	core "arkan-jaya/core/user"
)

//MySQLRepository The implementation of user.Repository object
type MySQLRepository struct {
	db *sql.DB
}

//NewMySQLRepository Generate mongo DB user repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

//FindUserByID Find user based on given ID. Its return nil if not found
func (repo *MySQLRepository) FindUserByID(ID string) (*core.User, error) {
	var user core.User

	selectQuery := `SELECT id, name, username, password, email, created_at, modified_at, role_id, is_active COALESCE(tags, "")
		FROM users i
		WHERE i.id = ?`

	err := repo.db.
		QueryRow(selectQuery, ID).
		Scan(
			&user.ID, &user.Name, &user.Username,
			&user.Email, &user.Password,
			&user.CreatedAt,
			&user.ModifiedAt,
			&user.RoleID, &user.IsActive)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

//FindAllUser Find all users based on given tag. Its return empty array if not found
func (repo *MySQLRepository) FindAllUser() ([]core.User, error) {
	//TODO: if feel have a performance issue in tag grouping, move the logic from db to here
	selectQuery := `SELECT id, name, username, password, email, created_at, modified_at, role_id, is_active COALESCE(tags, "")
	FROM users i`

	row, err := repo.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var users []core.User

	for row.Next() {
		var user core.User

		err := row.Scan(
			&user.ID, &user.Name, &user.Username,
			&user.Email, &user.Password,
			&user.CreatedAt,
			&user.ModifiedAt,
			&user.RoleID, &user.IsActive)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

//InsertUser Insert new user into database. Its return user id if success
func (repo *MySQLRepository) InsertUser(user core.User) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	userQuery := `INSERT INTO users (
			id, 
			name, 
			username, 
			email, 
			password, 
			created_at, 
			modified_at,
			role_id,
			is_active
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	if err != nil {
		return err
	}

	_, err = tx.Exec(userQuery,
		user.ID,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.ModifiedAt,
		user.RoleID,
		user.IsActive,
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

//UpdateUser Update existing user in database
func (repo *MySQLRepository) UpdateUser(user core.User) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	userInsertQuery := `UPDATE users
		SET
			name = ?,
			username = ?,
			email = ?,
			password = ?,
			modified_by = ?,
			role_id = ?
			is_active = ?
		WHERE id = ?`

	res, err := tx.Exec(userInsertQuery,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.ModifiedAt,
		user.RoleID,
		user.IsActive,
		user.ID,
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
