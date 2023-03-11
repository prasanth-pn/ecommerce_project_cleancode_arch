package handler

import (
	"clean/pkg/common/response"
	"clean/pkg/domain"
	"clean/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary List user for admin
// @ID AdminAddProducts
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param AdminAddProducts body domain.Product{} true "AdminAddProduct"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/add/products [post]
var p = fmt.Println

func (cr *AdminHandler) AddProducts(c *gin.Context) {
	//adding product details
	var products domain.Product
	file := c.PostForm("products")
	err := json.Unmarshal([]byte(file), &products)
	if err != nil {
		res := response.ErrorResponse("errror while json unmarshal", err.Error(), "error in addproduct")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	p(products.Description)
	if err != nil {
		res := response.ErrorResponse("error when binding products", err.Error(), "error in addproduct")
		c.Writer.WriteHeader(422)
		utils.ResponseJSON(c, res)
		return
	}
	product_id, err := cr.adminService.AddProducts(c.Request.Context(), products)
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
		err = c.SaveUploadedFile(imagepath, "./public"+image)
		if err != nil {
			res := response.ErrorResponse("error while savauploadfile", err.Error(), "error in saveupload")
			c.Writer.WriteHeader(422)
			utils.ResponseJSON(c, res)
			return
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
func (cr AdminHandler) DeleteProduct(c *gin.Context) {
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	p(product_id)
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
func (cr *AdminHandler) UpdateProduct(c *gin.Context) {
	product_id, _ := strconv.Atoi(c.Query("product_id"))
	fmt.Println(product_id)
	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		res := response.ErrorResponse("error while fetching data", err.Error(), "error while fetchign data in updateproduct")
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return
	}
	product.Product_Id = product_id
	err := cr.adminService.UpdateProduct(product)
	if err != nil {
		res := response.ErrorResponse("error while update product", err.Error(), product)
		c.Writer.WriteHeader(401)
		utils.ResponseJSON(c, res)
		return
	}
	res := response.SuccessResponse(true, "succefully updated", product)
	c.Writer.WriteHeader(401)
	utils.ResponseJSON(c, res)
}
func (cr *AdminHandler) ImageUpload(c *gin.Context) {
	//product_id,_:=strconv.Atoi(c.Query("product_id"))
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
		fmt.Println(image)
	}
	// fmt.Println(product_id)
	// err:=cr.adminService.ImageUpload(image,product_id)
	// if err!=nil{
	// 	res:=response.ErrorResponse("error from upload image",err.Error(),image)
	// }
}
