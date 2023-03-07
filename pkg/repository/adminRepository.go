package repository

import (
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"clean/pkg/utils"
	"database/sql"
)

type adminDatabase struct {
	DB *sql.DB
}

func NewAdminRepository(DB *sql.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}
func (c *adminDatabase) ListUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	var users []domain.UserResponse
	//	fmt.Println("reached int repository", pagenation.Page, pagenation.PageSize)
	query := `SELECT COUNT(*) OVER(),user_id,first_name,
last_name,
email,
gender,
phone
FROM users
LIMIT $1 OFFSET $2;`
	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}
	var totalRecords int
	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse
		err = rows.Scan(
			&totalRecords,
			&user.ID,
			&user.First_Name,
			&user.Last_Name,
			&user.Email,
			&user.Gender,
			&user.Phone,
		)
		if err != nil {
			return nil, utils.ComputeMetadata(&totalRecords, &pagenation.Page, &pagenation.PageSize), err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetadata(&totalRecords, &pagenation.Page, &pagenation.PageSize), err
	}
	return users, utils.ComputeMetadata(&totalRecords, &pagenation.Page, &pagenation.PageSize), err

}
func (c *adminDatabase) ListBlockedUsers(pagenation utils.Filter) ([]domain.Users, utils.Metadata, error) {
	var users []domain.Users
	query := `SELECT COUNT(*) OVER(), user_id,first_name,last_name,email,gender,phone  FROM users WHERE block_status=true LIMIT $1 OFFSET $2;`

	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	var totalrecords int
	defer rows.Close()
	for rows.Next() {
		var user domain.Users
		err = rows.Scan(&totalrecords, &user.User_Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Gender, &user.Phone)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
		users = append(users, user)
		if err := rows.Err(); err != nil {
			return users, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), err
		}
	}
	return users, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), err

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
			  category_id,brand_id)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`

	err := c.DB.QueryRow(query, product.Product_Name,
		product.Description,
		product.Quantity,
		product.Image_Path,
		product.Price,
		product.Color,
		product.Available,
		product.Trending,
		product.Category_Id,
		product.Brand_Id,
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

func (c *adminDatabase) AddBrand(brand domain.Brand) error {
	query := `INSERT INTO brands(brand_name,brand_description,discount)
	VALUES($1,$2,$3);`
	err := c.DB.QueryRow(query, brand.Brand_Name, brand.Brand_Description, brand.Discount).Err()

	return err

}

func (c *adminDatabase) AddModel(model domain.Model) error {
	query := `INSERT INTO models(model_name)VALUES($1);`
	err := c.DB.QueryRow(query, model.Model_Name).Err()
	return err
}
