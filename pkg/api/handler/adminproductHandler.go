package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	"clean/pkg/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary AddProducts for admin
// @ID AdminAddProducts
// @Tags PRODUCTMANAGEMENT
// @Produce json
// @Security BearerAuth
// @Param homepic formData file true "select the  image"
// @Param products formData string true "AdminAddProduct"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/add/products [post]
func (cr *AdminHandler) AddProducts(c *gin.Context) {
	//adding product details
	var products domain.Product
	file := c.PostForm("products")
	img, _ := c.FormFile("homepic")
	extention := filepath.Ext(img.Filename)
	imgf := "product" + uuid.New().String() + extention
	c.SaveUploadedFile(img, "./public/"+imgf)
	err := json.Unmarshal([]byte(file), &products)
	if err != nil {
		res := response.ErrorResponse("errror while json unmarshal", err.Error(), "error in addproduct")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	products.Image = imgf
	if err != nil {
		res := response.ErrorResponse("error when binding products", err.Error(), "error in addproduct")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	product_id, err := cr.adminService.AddProducts(products)
	products.Product_Id = product_id
	if err != nil {
		res := response.ErrorResponse("oops products not added", err.Error(), "proucts not added in add products some error")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(c, res)
		return
	}
	//inserting images
	var images []string
	if err = c.Request.ParseMultipartForm(32 << 20); err != nil {
		res := response.ErrorResponse("error while geting image", err.Error(), "error while parseimage")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	// get the image through parameter
	files := c.Request.MultipartForm.File["image"]
	for _, imagepath := range files {
		extention := filepath.Ext(imagepath.Filename)
		image := "product" + uuid.New().String() + extention
		err = c.SaveUploadedFile(imagepath, "./public/"+image)
		if err != nil {
			res := response.ErrorResponse("error while savauploadfile", err.Error(), "error in saveupload")
			c.Writer.WriteHeader(300)
			utils.ResponseJSON(c, res)
		}
		images = append(images, image)
	}
	err = cr.adminService.ImageUpload(images, product_id)
	if err != nil {
		res := response.ErrorResponse("error while image upload", err.Error(), "error in image upload service")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	respons := response.SuccessResponse(true, "SUCCESS", products, images)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(c, respons)
}

// @Summary AdminDeleteProducts for admin
// @ID AdminAdddProducts for admin
// @Tags PRODUCTMANAGEMENT
// @Produce json
// @Security BearerAuth
// @Param product_id query string true "product_id"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/product/delete [delete]
func (cr AdminHandler) DeleteProduct(c *gin.Context) {
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	product, err := cr.adminService.FindProduct(product_id)
	if err != nil {
		res := response.ErrorResponse("error whenn finding product", err.Error(), "error in find product in delete product")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	err = cr.adminService.DeleteProduct(product_id)
	if err != nil {
		res := response.ErrorResponse("error while delete the product", err.Error(), "error in delete products")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}

	res := response.SuccessResponse(true, "success", product)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary DeleteImagein Products
// @ID DeleteImageFromProduct for admin
// @Tags PRODUCTMANAGEMENT
// @Security BearerAuth
// @Param imageName query string true "select the image name"
// @Param product_id query string  true "enter the prouduct Id"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/products/delete-image [delete]
func (cr *AdminHandler) DeleteImage(c *gin.Context) {
	imageName := c.Query("imageName")
	product_id, _ := strconv.Atoi(c.Query("product_id"))

	err := cr.adminService.DeleteImage(product_id, imageName)
	fmt.Println(err)
	if err != nil {
		res := response.ErrorResponse("failed to delete the image ", err.Error(), "failed to delete the product")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	err = os.Remove("./public/" + imageName)
	if err != nil {
		res := response.ErrorResponse("failed to delete the image", err.Error(), "failed")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "success fully deleted", "the image"+imageName+" from product_id   :"+string(rune(product_id))+"deleted")
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary GenerateCoupon for admin
// @ID GenereateCoupon for admin
// @Tags COUPON
// @Produce json
// @Security BearerAuth
// @Param coupon query string false "coupon"
// @Param quantity query string true "quantity"
// @Param validity query string true "validity"
// @Param discount query  string true "discount"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/generate-coupon [post]
func (cr *AdminHandler) GenerateCoupon(c *gin.Context) {
	coupon := c.Query("coupon")
	quantity, _ := strconv.Atoi(c.Query("quantity"))
	validity, _ := strconv.Atoi(c.Query("validity"))
	discount, _ := strconv.Atoi(c.Query("discount"))
	expirationTime := time.Now().AddDate(0, 0, validity).Unix() //Add(2 * time.Minute).Unix() //
	if coupon == "" {
		length := 8
		source := rand.NewSource(time.Now().Unix())
		r := rand.New(source)
		charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		code := make([]byte, length)
		for i := range code {
			code[i] = charset[r.Intn(len(charset))]
		}
		coupon = string(code)
	}
	COUPON := domain.Coupon{
		Coupon:   coupon,
		Quantity: quantity,
		Validity: expirationTime,
		Discount: discount,
	}
	err := cr.adminService.GenerateCoupon(COUPON)
	if err != nil {
		res := response.ErrorResponse("failed to genereate coupon ", err.Error(), "failed to genarate coupon")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "coupon generated succefully", COUPON)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary ListCategories for admin
// @ID ListCategories for admin
// @Tags PRODUCTMANAGEMENT
// @Product json
// @Security BearerAuth
// @Param page query string true "page"
// @Param pagesize query string true "pagesize"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/list/category [get]
func (cr *AdminHandler) ListCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	categories, metadata, err := cr.adminService.ListCategories(pagenation)
	if err != nil {
		res := response.ErrorResponse("error from listcategories", err.Error(), "error in list categories")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "success", categories, metadata)
	c.Writer.WriteHeader(422)
	utils.ResponseJSON(c, res)
}

// @Summary ListProductByCategories for admin
// @ID listproductsbycategories for admin
// @Tags PRODUCTMANGEMENT
// @Product json
// @Security BearerAuth
// @Param cat_id query string true "category_id"
// @Param page query string true "page"
// @Param pagesize query string true "pagesize"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/list/productby-categories [get]
func (cr *AdminHandler) ListProductsByCategories(c *gin.Context) {
	category_id, _ := strconv.Atoi(c.Query("cat_id"))
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenation := utils.Filter{
		Page:     page,
		PageSize: pagesize,
	}
	products, metadata, err := cr.adminService.ListProductByCategories(pagenation, category_id)

	if err != nil {
		res := response.ErrorResponse("error in list productsby categories", err.Error(), "error in list product by categories")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "succefully listed products by categories", products, metadata)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

// @Summary UpdateProduct for admin
// @ID UpdateProduct for admin
// @Tags PRODUCTMANAGEMENT
// @Product json
// @Security BearerAuth
// @Param product_id query string true "product_id"
// @Param image formData file true "select image"
// @Param updateproduct formData string true "updateproduct"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/update/product [patch]
func (cr *AdminHandler) UpdateProduct(c *gin.Context) {
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	updateproduct := c.PostForm("updateproduct")
	file, err := c.FormFile("image")
	if err != nil {
		res := response.ErrorResponse("please select a image file ", err.Error(), "select a image file to upload")
		c.Writer.WriteHeader(300)
		utils.ResponseJSON(c, res)
	}
	extention := filepath.Ext(file.Filename)
	image := "product" + uuid.New().String() + extention
	fmt.Println(image)
	var product domain.Product
	if err := json.Unmarshal([]byte(updateproduct), &product); err != nil {
		res := response.ErrorResponse("error while fetching data", err.Error(), "error while fetchign data in updateproduct")
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return
	}
	fmt.Println(product.Image)
	product.Image = image
	product.Product_Id = product_id
	fmt.Println(product)
	err = cr.adminService.UpdateProduct(product)
	if err != nil {
		res := response.ErrorResponse("error while update product", err.Error(), product)
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return
	}
	c.SaveUploadedFile(file, "./public/"+image)
	res := response.SuccessResponse(true, "succefully updated", product)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}

//	@Summary ImageUploadproduct for admin
//	@ID productImageUpload for admin
//
// @Tags PRODUCTMANAGEMENT
// @Product json
// @Security BearerAuth
// @Param product_id query string true "product_id"
// @Param image formData file true "select image"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /admin/upload/image [patch]
func (cr *AdminHandler) ImageUpload(c *gin.Context) {
	var images []string
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		res := response.ErrorResponse("error getting images", err.Error(), "error getting images")
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return
	}
	file := c.Request.MultipartForm.File["image"]
	for _, imagepath := range file {
		extention := filepath.Ext(imagepath.Filename)
		fmt.Println(extention)
		image := "product" + uuid.New().String() + extention
		images = append(images, image)
	}
	fmt.Println(product_id)
	err := cr.adminService.ImageUpload(images, product_id)
	if err != nil {
		res := response.ErrorResponse("error from upload image", err.Error(), "image isn not uploaded")
		c.Writer.WriteHeader(400)
		utils.ResponseJSON(c, res)
		return
	}
	data := struct {
		Image      []string
		Product_id int
	}{
		Image:      images,
		Product_id: product_id,
	}
	res := response.SuccessResponse(true, "successfully image uploaded", data)
	c.Writer.WriteHeader(200)
	utils.ResponseJSON(c, res)
}
