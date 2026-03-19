package pantry

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(userID string) ([]PantryItem, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, name, quantity, unit, created_at, updated_at FROM pantry_items WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PantryItem
	for rows.Next() {
		var item PantryItem
		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.Name,
			&item.Quantity,
			&item.Unit,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) Create(item *PantryItem) error {
	err := r.db.QueryRow(
		`INSERT INTO pantry_items (user_id, name, quantity, unit)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`,
		item.UserID,
		item.Name,
		item.Quantity,
		item.Unit,
	).Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt)

	return err
}

func (r *Repository) Update(item *PantryItem) error {
	err := r.db.QueryRow(
		`UPDATE pantry_items 
		SET name = $1, quantity = $2, unit = $3, updated_at = NOW()
		WHERE id = $4 AND user_id = $5
		RETURNING created_at, updated_at`,
		item.Name,
		item.Quantity,
		item.Unit,
		item.ID,
		item.UserID,
	).Scan(&item.CreatedAt, &item.UpdatedAt)

	return err
}

func (r *Repository) Delete(item *PantryItem) error {
	_, err := r.db.Exec(
		`DELETE FROM pantry_items 
		WHERE id = $1 AND user_id = $2`,
		item.ID,
		item.UserID,
	)
	return err
}
