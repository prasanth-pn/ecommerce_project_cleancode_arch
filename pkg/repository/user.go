package repository

import (
	//domain "clean/pkg/domain"
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"errors"
	"fmt"

	"database/sql"
	//"gorm.io/gorm"
)

type userDatabase struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
func (c *userDatabase) ListProducts() ([]domain.ProductResponse, error) {
	var products []domain.ProductResponse
	query := `SELECT P.product_id,P.product_name,P.description,P.image_path,P.price,P.color,P.available,P.trending,
C.category_name FROM products AS P
INNER JOIN categories AS C
ON P.product_id=C.category_id;`
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, errors.New("some mistake in query")
	}
	defer rows.Close()
	for rows.Next() {
		var product domain.ProductResponse
		err := rows.Scan(
			&product.Product_Id,
			&product.Product_Name,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.Color,
			&product.Available,
			&product.Trending,
			&product.Category_Name,
		)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while scan database")
		}
		products = append(products, product)

	}
	return products, nil
}
func (c *userDatabase) FindProduct(product_id uint) (domain.Product, error) {
	var Product domain.Product
	query := `SELECT price,quantity FROM products
	WHERE product_id=$1;`
	err := c.DB.QueryRow(query, product_id).Scan(&Product.Price,
		&Product.Quantity,
	)
	return Product, err
}
func (c *userDatabase) ListCart(User_id uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	query := `SELECT cart_id,product_id FROM carts WHERE user_id=$1;`

	rows, err := c.DB.Query(query, User_id)
	if err != nil {
		return carts, errors.New("error in query")

	}
	defer rows.Close()
	for rows.Next() {
		var cart domain.Cart
		err = rows.Scan(&cart.Cart_id,
			&cart.ProductID)

		if err != nil {
			return carts, errors.New("error while scaning in database")
		}
		carts = append(carts, cart)

	}
	return carts, nil
}

func (c *userDatabase) QuantityCart(product_id, user_id uint) (domain.Cart, error) {
	var cart domain.Cart
	query := `SELECT quantity FROM carts WHERE product_id=$1 and user_id=$2;`
	err := c.DB.QueryRow(query, product_id, user_id).Scan(&cart.Quantity)
	return cart, err
}
func (c *userDatabase) UpdateCart(totalprice float32, quantity, product_id, user_id uint) error {
	var cart domain.Cart
	query := `UPDATE carts SET quantity=$1,total_price=$2 WHERE product_id=$3 and user_id=$4;`

	err := c.DB.QueryRow(query, quantity, totalprice, product_id, user_id).Scan(&cart.Quantity,
		&cart.Total_Price)
	return err

}

func (c *userDatabase) CreateCart(cart domain.Cart) error {
	query := `INSERT INTO carts(user_id,product_id,quantity,total_price)
	values($1,$2,$3,$4);`
	err := c.DB.QueryRow(query, cart.User_Id, cart.ProductID, cart.Quantity, cart.Total_Price).Err()
	return err
}

// i.DB.Raw("update carts set quantity=?,total_price=? where product_id=? and user_id=? ", prodqua+cart.Quantity, totl, prodid, user.ID).Scan(&Cart)
