package repository

import (
	//domain "clean/pkg/domain"
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	"clean/pkg/utils"
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
func (c *userDatabase) ListProducts(pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse
	query := `SELECT COUNT(*) OVER(), P.product_id,P.product_name,P.description,P.image_path,P.price,P.color,P.available,P.trending,
C.category_name FROM products AS P
INNER JOIN categories AS C
ON P.category_id=C.category_id
where P.category_id=3
LIMIT $1 OFFSET $2;`
	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, errors.New("some mistake in query")
	}
	var totalrecords int
	defer rows.Close()
	for rows.Next() {
		var product domain.ProductResponse
		err := rows.Scan(
			&totalrecords,
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
			return nil, utils.Metadata{}, errors.New("error while scan database")
		}
		products = append(products, product)

	}
	return products, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), err
}
func (c *userDatabase) ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error) {
	var categories []domain.Category
	query := `SELECT COUNT(*) OVER() ,category_id,category_name,description,image FROM categories LIMIT $1 OFFSET $2;`
	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	var totalrows int
	defer rows.Close()
	var catego domain.Category
	for rows.Next() {
		err := rows.Scan(&totalrows, &catego.Category_Id, &catego.Category_Name, &catego.Description, &catego.Image)

		if err != nil {
			return nil, utils.Metadata{}, err
		}
		categories = append(categories, catego)
	}
	return categories, utils.ComputeMetadata(&totalrows, &pagenation.Page, &pagenation.PageSize), err
}

func (c *userDatabase) ListProductsByCategories(category_id int, pagenation utils.Filter) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse
	var Images []string
	query := `SELECT COUNT(*) OVER(), P.product_id,P.product_name,P.description,P.image,P.price,
	C.category_name FROM products AS P INNER JOIN categories AS C ON P.category_id=C.category_id  WHERE P.category_id=$1 LIMIT $2 OFFSET $3;`
	imgquery := `SELECT image FROM images WHERE product_id=$1;`

	rows, err := c.DB.Query(query, category_id, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	defer rows.Close()
	var totalrecords int
	for rows.Next() {
		var Image string
		var product domain.ProductResponse
		err = rows.Scan(&totalrecords, &product.Product_Id, &product.Product_Name, &product.Description, &product.MainPic, &product.Price, &product.Category_Name)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
		img, err := c.DB.Query(imgquery, product.Product_Id)
		if err != nil {
			return products, utils.Metadata{}, err
		}
		defer img.Close()
		for img.Next() {
			err = img.Scan(&Image)
			if err != nil {
				return products, utils.Metadata{}, err
			}
			Images = append(Images, Image)

		}
		product.Image = Images

		products = append(products, product)
		Images = nil

	}
	return products, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), err

}
func (c *userDatabase) FindProduct(product_id uint) (domain.Product, error) {
	var Product domain.Product
	query := `SELECT * FROM products
	WHERE product_id=$1;`
	err := c.DB.QueryRow(query, product_id).Scan(&Product.Product_Id, &Product.Product_Name, &Product.Image, &Product.Description, &Product.Quantity, &Product.Price,
		&Product.Color, &Product.Available, &Product.Trending, &Product.Category_Id, &Product.Brand_Id,
	)
	return Product, err
}
func (c *userDatabase) ListCart(pagenation utils.Filter, User_id uint) ([]domain.CartListResponse, utils.Metadata, error) {
	var carts []domain.CartListResponse
	var totalrecords int
	query := `SELECT COUNT(*) OVER() ,C.created_at,C.quantity,P.price,C.total_price,P.description,P.image,P.product_name,C.product_id
	FROM products AS P
	INNER JOIN carts AS C
	ON C.product_id=P.product_id
	WHERE C.user_id=$1 LIMIT $2 OFFSET  $3;`
	rows, err := c.DB.Query(query, User_id, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var cart domain.CartListResponse
		err = rows.Scan(&totalrecords, &cart.Created_At, &cart.Quantity, &cart.Price, &cart.Total_Amount, &cart.Description, &cart.Image_Path, &cart.Product_Name, &cart.Product_id)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
		carts = append(carts, cart)

	}
	return carts, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), nil
}

func (c *userDatabase) QuantityCart(product_id, user_id uint) (domain.Cart, error) {
	var cart domain.Cart
	query := `SELECT quantity FROM carts WHERE product_id=$1 and user_id=$2;`
	err := c.DB.QueryRow(query, product_id, user_id).Scan(&cart.Quantity)
	return cart, err
}
func (c *userDatabase) UpdateCart(totalprice float32, quantity, product_id, user_id uint) (domain.Cart, error) {
	var time = time.Now()
	var Cart domain.Cart
	query := `UPDATE carts SET  quantity=$1,total_price=$2 ,updated_at=$3 WHERE product_id=$4 and user_id=$5 RETURNING created_at,user_id,image,product_id,quantity,total_price;`
	err := c.DB.QueryRow(query, quantity, totalprice, time, product_id, user_id).Scan(&Cart.Created_At, &Cart.User_Id, &Cart.Image, &Cart.Product_Id,
		&Cart.Quantity, &Cart.Total_Price)
	return Cart, err
}

func (c *userDatabase) CreateCart(cart domain.Cart) (domain.Cart, error) {
	var time = time.Now()
	var Cart domain.Cart
	query := `INSERT INTO carts(created_at,user_id,image,product_id,quantity,total_price)
	values($1,$2,$3,$4,$5,$6) RETURNING created_at,user_id,image,product_id,quantity,total_price;`
	err := c.DB.QueryRow(query, time, cart.User_Id, cart.Image, cart.Product_Id, cart.Quantity, cart.Total_Price).Scan(&Cart.Created_At,
		&Cart.User_Id, &Cart.Image, &Cart.Product_Id, &Cart.Quantity, &Cart.Total_Price)
	return Cart, err
}
func (c *userDatabase) ListViewCart(user_id uint) ([]domain.CartListResponse, error) {
	var carts []domain.CartListResponse
	query := `SELECT C.quantity,C.total_price,P.description,P.image,P.product_name,C.product_id
	FROM products AS P
	INNER JOIN carts AS C
	ON C.product_id=P.product_id
	WHERE C.user_id=$1;`
	rows, err := c.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cart domain.CartListResponse
	for rows.Next() {
		err := rows.Scan(&cart.Quantity,
			&cart.Total_Amount,
			&cart.Description,
			&cart.Image_Path,
			&cart.Product_Name,
			&cart.Product_id,
		)
		if err != nil {
			return nil, err
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
	//defer rows.Close()
	return value, err
}
func (c *userDatabase) FindTheSumOfCart(user_id int) (int, error) {
	var sum int
	query := `SELECT SUM(total_price) FROM carts WHERE user_id=$1;`
	err := c.DB.QueryRow(query, user_id).Scan(&sum)
	return sum, err
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
	return err
}
func (c *userDatabase) ViewWishList(user_id uint) []domain.WishListResponse {
	var wish []domain.WishListResponse
	query := `select wish_lists.id,users.user_id,products.product_id,products.product_name ,products.price,products.image from wish_lists join products 
	on wish_lists.product_id=products.product_id 
	join users on wish_lists.user_id=users.user_id where users.user_id=$1;`
	var wishe domain.WishListResponse
	rows, err := c.DB.Query(query, user_id)
	fmt.Println(err)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&wishe.WishList_Id, &wishe.User_id,
			&wishe.Product_id,
			&wishe.Product_name,
			&wishe.Price, &wishe.Image)
		wish = append(wish, wishe)
	}
	return wish
}
func (c *userDatabase) RemoveFromWishlist(user_id, product_id int) error {
	query := `DELETE FROM wish_lists 	WHERE user_id=$1 AND product_id=$2;`
	err := c.DB.QueryRow(query, user_id, product_id)
	return err.Err()
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
func (c *userDatabase) FindAddress(user_id, address_id uint) (domain.Address, error) {
	var address domain.Address
	query := `SELECT * FROM addresses WHERE user_id=$1 AND address_id=$2;`
	err := c.DB.QueryRow(query, user_id, address_id).Scan(&address.Address_id,
		&address.User_Id,
		&address.FName,
		&address.LName,
		&address.Phone_Number,
		&address.Pincode,
		&address.House,
		&address.Area,
		&address.Landmark,
		&address.City)
	return address, err
}
func (c *userDatabase) UpdateAddress(add domain.Address, user_id, address_id uint) error {
	query := `UPDATE addresses SET l_name=$1 ,f_name=$2,phone_number=$3,pincode=$4,house=$5,area=$6,landmark=$7,city=$8 WHERE user_id=$9 AND address_id=$10;`
	err := c.DB.QueryRow(query, add.LName, add.FName, add.Phone_Number, add.Pincode, add.House, add.Area, add.Landmark, add.City, user_id, address_id).Err()
	fmt.Println(err, "repository error if ")
	return err
}
func (c *userDatabase) CreateOrder(order domain.Orders) error {
	var time = time.Now()
	query := `INSERT INTO orders (created_at,user_id,order_id,total_amount,payment_method,payment_status,order_status,address_id) values($1,$2,$3,$4,$5,$6,$7,$8);`
	err := c.DB.QueryRow(query, time, order.User_Id, order.Order_Id, order.Total_Amount, order.PaymentMethod, order.Payment_Status, order.Order_Status, order.Address_Id)
	return err.Err()
}
func (c *userDatabase) SearchOrder(order_id string) (domain.Orders, error) {
	var order domain.Orders
	query := `SELECT created_at,user_id,total_amount,payment_method,payment_status,order_status,address_id FROM orders where order_id=$1;`
	err := c.DB.QueryRow(query, order_id).Scan(&order.Created_at, &order.User_Id, &order.Total_Amount, &order.PaymentMethod, &order.Payment_Status, &order.Order_Status, &order.Address_Id)
	return order, err
}
func (c *userDatabase) UpdateOrders(payement_id, order_id string) error {
	query := `UPDATE orders SET payment_status=$1,payment_id=$2 WHERE order_id=$3;`
	err := c.DB.QueryRow(query, "payment successful", payement_id, order_id).Err()
	return err
}
func (c *userDatabase) Insert_To_My_Order(carts domain.CartListResponse, order_id string) error {
	query := `INSERT INTO orderd_items(product_id,order_id,product_name,image_path,price,quantity)
	 values($1,$2,$3,$4,$5,$6);`
	err := c.DB.QueryRow(query, carts.Product_id, order_id, carts.Product_Name, carts.Image_Path, carts.Total_Amount, carts.Quantity)
	return err.Err()
}
func (c *userDatabase) ClearCart(user_id uint) error {
	query := `DELETE FROM carts WHERE user_id=$1;`
	err := c.DB.QueryRow(query, user_id)
	return err.Err()
}
func (c *userDatabase) ListOrder(user_id uint) ([]domain.ListOrder, uint, error) {
	var orders []domain.ListOrder
	query := `select P.created_at ,P.order_id,P.total_amount,P.payment_method,P.payment_status,P.payment_id,P.order_status,P.address_id,C.product_id,
	C.product_name,C.image_path,C.price,C.quantity from orders AS P
	inner join orderd_items AS C ON C.order_id=P.order_id WHERE user_id=$1;`
	rows, err := c.DB.Query(query, user_id)
	if err != nil {
		return nil, 0, err
	}
	var add_id uint
	var order domain.ListOrder
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&order.Created_At, &order.Order_id, &order.Total_amount,
			&order.Payment_Method, &order.Payment_Status, &order.Payment_id, &order.Order_Status, &order.Address_Id,
			&order.Product_Id, &order.Product_Name, &order.Image_Path, &order.Price, &order.Quantity)
		if err != nil {
			return nil, 0, err
		}
		orders = append(orders, order)
		add_id = order.Address_Id
	}
	return orders, add_id, nil
}
func (c *userDatabase) FindCoupon(coupon string) (domain.Coupon, error) {
	var Coupon domain.Coupon
	query := `SELECT * FROM coupons WHERE coupon=$1;`
	err := c.DB.QueryRow(query, coupon).Scan(&Coupon.Created_At, &Coupon.Coupon_Id, &Coupon.Coupon, &Coupon.Discount, &Coupon.Quantity, &Coupon.Validity)
	return Coupon, err
}
