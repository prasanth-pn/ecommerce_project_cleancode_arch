package repository

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	interfaces "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/repository/interfaces"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
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
	query := `SELECT COUNT(*) OVER(), user_id,first_name,last_name,email,gender,phone,country,city FROM users WHERE block_status=$1 LIMIT $2 OFFSET $3;`

	rows, err := c.DB.Query(query, true, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	var totalrecords int
	defer rows.Close()
	for rows.Next() {
		var user domain.Users
		err = rows.Scan(&totalrecords, &user.User_Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Gender, &user.Phone, &user.Country, &user.City)
		fmt.Println(err, "error in next")
		if err != nil {
			return nil, utils.Metadata{}, err
		}
		users = append(users, user)
	}
	return users, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), err
}
func (c *adminDatabase) AddProducts(product domain.Product) (int, error) {
	query := `INSERT INTO 
	products (product_name,
			  image,
			  description,
			  quantity,
			  price,
			  color,
			  available,
			  trending,
			  category_id,brand_id)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)RETURNING product_id;`
	var product_id int
	err := c.DB.QueryRow(query, product.Product_Name, product.Image,
		product.Description,
		product.Quantity,
		product.Price,
		product.Color,
		product.Available,
		product.Trending,
		product.Category_Id,
		product.Brand_Id,
	).Scan(&product_id)
	return product_id, err
}
func (c *adminDatabase) FindProduct(product_id int) (domain.ProductResponse, error) {
	var products domain.ProductResponse
	var images []string
	query := `SELECT P.product_id,P.product_name,P.description,P.price,P.color,P.available,P.trending,B.brand_name,C.category_name,I.image
	FROM products AS P INNER JOIN brands AS B
	ON P.brand_id=B.brand_id
	LEFT JOIN categories AS C
	ON P.category_id=C.category_id INNER JOIN images AS I
    ON P.product_id=I.product_id WHERE P.product_id =$1;`
	rows, err := c.DB.Query(query, product_id)
	if err != nil {

		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		var product domain.ProductResponse
		var image string
		err = rows.Scan(&product.Product_Id, &product.Product_Name, &product.Description, &product.Price, &product.Color, &product.Available, &product.Trending, &product.Brand_Name, &product.Category_Name, &image)
		if err != nil {
			fmt.Println("first error")
			return products, err
		}

		products = product
		images = append(images, image)
	}
	products.Image = images

	return products, nil
}
func (c *adminDatabase) DeleteProduct(product_id int) error {
	//product_id=2
	query := `DELETE FROM products WHERE product_id=$1;`
	err := c.DB.QueryRow(query, product_id).Err()
	fmt.Println(err)
	if err != nil {
		return err
	}
	que := `DELETE FROM images where product_id=$1;`
	err = c.DB.QueryRow(que, product_id).Err()
	if err != nil {
		return err
	}
	return nil
}
func (c *adminDatabase) AddCategory(category domain.Category) error {
	query := `INSERT INTO categories(category_name,description,image)
	VALUES($1,$2,$3);`
	err := c.DB.QueryRow(query, category.Category_Name,
		category.Description, category.Image).Err()
	return err
}
func (c *adminDatabase) ListCategories(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error) {
	var categories []domain.Category
	var totalrecords int
	query := `SELECT COUNT(*) OVER(),* FROM categories LIMIT $1 OFFSET $2;`
	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return categories, utils.Metadata{}, err
	}
	defer rows.Close()
	for rows.Next() {

		var category domain.Category
		err = rows.Scan(&totalrecords, &category.Category_Id, &category.Category_Name, &category.Description, &category.Image)
		if err != nil {
			return categories, utils.Metadata{}, err
		}
		categories = append(categories, category)
	}
	return categories, (utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize)), err
}
func (c *adminDatabase) ListProductByCategories(pagenation utils.Filter, cate_id int) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse
	var totalrecords int
	var images []string
	query := `SELECT COUNT(*) OVER() ,P.product_id,P.product_name,P.description,P.image,P.price,P.color,P.available,P.trending,C.category_name,B.brand_name
	FROM products AS P INNER JOIN categories AS C ON P.category_id=C.category_id INNER JOIN brands AS B ON P.brand_id=B.brand_id WHERE C.category_id=$1 LIMIT $2 OFFSET $3;`
	img := `SELECT image FROM images WHERE product_id=$1;`
	rows, err := c.DB.Query(query, cate_id, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return products, utils.Metadata{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var image string
		var product domain.ProductResponse
		err = rows.Scan(&totalrecords, &product.Product_Id, &product.Product_Name, &product.Description, &product.MainPic, &product.Price, &product.Color, &product.Available, &product.Trending, &product.Category_Name,
			&product.Brand_Name)
		if err != nil {
			return products, utils.Metadata{}, err
		}
		fmt.Println(product.Product_Id)
		irows, err := c.DB.Query(img, product.Product_Id)
		if err != nil {
			return products, utils.Metadata{}, err
		}
		defer irows.Close()

		for irows.Next() {
			irows.Scan(&image)

			images = append(images, image)
		}
		product.Image = images
		products = append(products, product)
		images = nil

	}
	fmt.Println(images)
	return products, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), err
}
func (c *adminDatabase) UpdateProduct(product domain.Product) error {
	fmt.Println(product, "in repository")
	query := `UPDATE products SET product_name=$1,description=$2,quantity=$3,price=$4,color=$5,available=$6,trending=$7, category_id=$8,
	brand_id=$9,image=$10 WHERE product_id=$11;`
	err := c.DB.QueryRow(query, product.Product_Name, product.Description, product.Quantity, product.Price, product.Color, product.Available, product.Trending,
		product.Category_Id, product.Brand_Id, product.Image, product.Product_Id)
	return err.Err()
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
func (c *adminDatabase) ImageUpload(image []string, product_id int) error {
	que := `CREATE TABLE IF NOT EXISTS images(
		id SERIAL PRIMARY KEY,
		product_id integer NOT NULL,
		image TEXT
	);`
	_, err := c.DB.Exec(que)
	if err != nil {
		return err
	}
	//images:=pq.Array(image)
	// images := pq.StringArray(image)
	fmt.Println(err)
	rows, err := c.DB.Prepare("INSERT INTO images (product_id, image) VALUES ($1, $2);")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	// Execute the insert statement
	for _, img := range image {
		_, err = rows.Exec(product_id, (img))
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	// imageValue := "productf2a4b318-79c0-416a-8883-6b07a81e80e0.jpeg"
	// _, err = c.DB.Exec("UPDATE images SET image = ARRAY(SELECT * FROM unnest(image) WHERE unnest != $1) WHERE id = $2", imageValue, 1)
	fmt.Println("Data inserted successfully!")
	return nil
}
func (c *adminDatabase) DeleteImage(product_id int, imagename string) error {
	query := `DELETE from images where product_id=$1 AND image=$2;`
	err := c.DB.QueryRow(query, product_id, imagename)
	return err.Err()
}
func (c *adminDatabase) GenerateCoupon(coupon domain.Coupon) error {
	var time = time.Now()
	query := `INSERT INTO coupons (created_at,coupon,discount,quantity,validity)VALUES($1,$2,$3,$4,$5);`
	err := c.DB.QueryRow(query, time, coupon.Coupon, coupon.Discount, coupon.Quantity, coupon.Validity)
	return err.Err()
}
func (c *adminDatabase) FindCoupon(coupon string) (domain.Coupon, error) {
	var cpn domain.Coupon
	query := `SELECT * FROM coupons where coupon=$1;`
	err := c.DB.QueryRow(query, coupon).Scan(&cpn.Created_At, &cpn.Coupon_Id, &cpn.Coupon, &cpn.Discount, cpn.Quantity, cpn.Validity)
	return cpn, err
}
func (c *adminDatabase) SearchUserByName(pagenation utils.Filter, name string) ([]domain.Users, utils.Metadata, error) {
	var users []domain.Users
	var totalrecords int
	query := `SELECT COUNT(*) OVER(),user_id,first_name,last_name,country,city,phone,email,gender FROM users WHERE first_name LIKE $1 LIMIT $2 OFFSET $3;`
	rows, err := c.DB.Query(query, name, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var user domain.Users
		err := rows.Scan(&totalrecords, &user.User_Id, &user.First_Name, &user.Last_Name, &user.Country, &user.City, &user.Phone, &user.Email, &user.Gender)
		if err != nil {
			return users, utils.Metadata{}, err
		}
		users = append(users, user)
	}
	return users, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), nil
}
func (c *adminDatabase) ListOrder(pagenation utils.Filter) ([]domain.Orders, utils.Metadata, error) {
	var Orders []domain.Orders
	var totalrecords int
	query := `SELECT COUNT(*) OVER(),* FROM orders LIMIT $1 OFFSET $2;`
	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return Orders, utils.Metadata{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var Canceled_at []byte

		var order domain.Orders
		err = rows.Scan(&totalrecords, &order.Created_at, &Canceled_at, &order.User_Id, &order.Order_Id, &order.Applied_Coupons,
			&order.Discount, &order.Total_Amount, &order.Balance_Amount, &order.PaymentMethod, &order.Payment_Status,
			&order.Payment_Id, &order.Order_Status, &order.Address_Id)
		if err != nil {
			return Orders, utils.Metadata{}, err
		}
		Orders = append(Orders, order)
	}
	return Orders, utils.ComputeMetadata(&totalrecords, &pagenation.Page, &pagenation.PageSize), nil
}
