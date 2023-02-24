package repository

import (
	//domain "clean/pkg/domain"
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"database/sql"
	"errors"
	"fmt"
	"time"
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
		fmt.Println(product)
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
		return nil, errors.New("error in query")

	}
	defer rows.Close()
	for rows.Next() {
		var cart domain.Cart
		err = rows.Scan(&cart.Cart_Id,
			&cart.Product_Id)

		if err != nil {
			return nil, errors.New("error while scaning in database")
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
	var time = time.Now()
	query := `UPDATE carts SET  quantity=$1,total_price=$2 ,updated_at=$3 WHERE product_id=$4 and user_id=$5 ;`

	err := c.DB.QueryRow(query, quantity, totalprice, time, product_id, user_id).Scan(&cart.Quantity,
		&cart.Total_Price)
	fmt.Println(err)
	return err

}

func (c *userDatabase) CreateCart(cart domain.Cart) error {
	var time = time.Now()
	query := `INSERT INTO carts(created_at,user_id,product_id,quantity,total_price)
	values($1,$2,$3,$4,$5);`
	err := c.DB.QueryRow(query, time, cart.User_Id, cart.Product_Id, cart.Quantity, cart.Total_Price).Err()
	return err
}

// i.DB.Raw("update carts set quantity=?,total_price=? where product_id=? and user_id=? ", prodqua+cart.Quantity, totl, prodid, user.ID).Scan(&Cart)

func (c *userDatabase) ViewCart(user_id uint) ([]domain.CartListResponse, error) {
	var carts []domain.CartListResponse

	query := `SELECT C.quantity,C.total_price,P.description,P.image_path,P.product_name
	FROM products AS P
	INNER JOIN carts AS C
	ON C.cart_id=P.product_id
	WHERE C.user_id=$1;`
	rows, err := c.DB.Query(query, user_id)
	fmt.Println(err, "error in repo")

	if err != nil {
		return nil, errors.New("error is happend in viewcart while query")
	}
	defer rows.Close()
	var cart domain.CartListResponse
	for rows.Next() {

		err := rows.Scan(&cart.Quantity,
			&cart.Total_Price,
			&cart.Description,
			&cart.Image_Path,
			&cart.Product_Name,
		)
		if err != nil {
			return nil, errors.New("errors while returning  values to cart response")
		}
		carts = append(carts, cart)

	}
	return carts, nil
}
func (c *userDatabase) TotalCartPrice(user_id uint) (float32, error) {
	query := `SELECT sum(total_price) from carts
	WHERE user_id=$1;`
	var value float32
	err := c.DB.QueryRow(query, user_id).Scan(&value)
	if err != nil {
		return value, errors.New("errror in total casrtsprice repo")
	}
	//defer rows.Close()
	return value, nil
}
func (c *userDatabase) Count_WishListed_Product(user_id, product_id uint) int {
	var count int

	query := `SELECT COUNT(*) FROM wish_lists WHERE user_id=$1 AND product_id=$2;`

	err := c.DB.QueryRow(query, user_id, product_id).Scan(&count)
	if err != nil {
		fmt.Println("error in query")
		return 0
	}
	return count
}
func (c *userDatabase) AddTo_WishList(wishlist domain.WishList) error {
	query := `INSERT INTO wish_lists(user_id , product_id)VALUES($1,$2);`
	err := c.DB.QueryRow(query, wishlist.UserID, wishlist.Product_Id).Err()
	if err != nil {
		return err
	}
	fmt.Println(err)

	return nil
}
func (c *userDatabase) ViewWishList(user_id uint) []domain.WishListResponse {
	var wish []domain.WishListResponse
	query := `select users.user_id,products.product_id,products.product_name ,products.price from wish_lists join products 
	on wish_lists.product_id=products.product_id 
	join users on wish_lists.user_id=users.user_id where users.user_id=$1;`
	var wishe domain.WishListResponse
	rows, err := c.DB.Query(query, user_id)
	fmt.Println(err)

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&wishe.Product_id,
			&wishe.User_id,
			&wishe.Product_name,
			&wishe.Price)
		wish = append(wish, wishe)
	}
	return wish
}
func (c *userDatabase) RemoveFromWishlist(user_id, product_id int) {
	query := `DELETE FROM wish_lists 	WHERE user_id=$1 AND product_id=$2;`
	err := c.DB.QueryRow(query, user_id, product_id)
	if err != nil {
		fmt.Println(" this is not deletd", err)
	}

}
func (c *userDatabase) FindCart(user_id, product_id uint) (domain.CartResponse, error) {
	var cart domain.CartResponse

	query := `SELECT quantity,total_price FROM carts WHERE user_id=$1 AND product_id=$2;`
	err := c.DB.QueryRow(query, user_id, product_id).Scan(&cart.Quantity,
		&cart.Total_Price,
	)
	fmt.Println(err, "this is the error")
	return cart, err

}


func(c *userDatabase)FindAddress(user_id,address_id uint)(domain.Address,error){
	var address domain.Address

	query :=`SELECT * FROM 	addresses WHERE user_id=$1 AND address_id=$2;`

	err:=c.DB.QueryRow(query,user_id,address_id).Scan(&address.Address_id,
	&address.Area,
&address.City,
&address.FName,
&address.House,
&address.LName,
address.Landmark,
&address.Pincode)

return address,err

}
