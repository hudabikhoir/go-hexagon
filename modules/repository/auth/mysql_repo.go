package auth

import (
	"database/sql"

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

//FindUserByUsername Find user based on given username. Its return nil if not found
func (repo *MySQLRepository) FindUserByUsername(username string) (*core.User, error) {
	var user core.User

	selectQuery := `SELECT id, name, username, password, email, created_at, modified_at, role_id, is_active COALESCE(tags, "")
		FROM users i
		WHERE i.username = ?`

	err := repo.db.
		QueryRow(selectQuery, username).
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
