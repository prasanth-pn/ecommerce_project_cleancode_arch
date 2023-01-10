package repository

import (
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"database/sql"
)

type adminDatabase struct {
	DB *sql.DB
}

func NewAdminRepository(DB *sql.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}
func (c *adminDatabase) ListUsers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT user_id,first_name,
last_name,
email,
gender,
phone
FROM users;`
	rows, err := c.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.First_Name,
			&user.Last_Name,
			&user.Email,
			&user.Gender,
			&user.Phone,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil

}
func (c *adminDatabase) AddProducts(product domain.Product) error {
	query := `INSERT INTO 
	products (product_name,
			  description,
			  quantity,
			  image_path,
			  price,
			  color,
			  available,
			  trending,
			  category_id)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);`

	err := c.DB.QueryRow(query, product.Product_name,
		product.Description,
		product.Quantity,
		product.Image_Path,
		product.Price,
		product.Color,
		product.Available,
		product.Trending,
		product.Category_id,
	).Err()
	return err
}
func (c *adminDatabase) AddCategory(category domain.Category) error {
	query := `INSERT INTO categories(category_name,description,image)
	VALUES($1,$2,$3);`
	err := c.DB.QueryRow(query, category.Category_Name,
		category.Description, category.Image).Err()
	return err
}
