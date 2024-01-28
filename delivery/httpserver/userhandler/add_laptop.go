package userhandler

import (
	"errors"
	"fmt"
	"game-app/param"
	"game-app/pkg/claim"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func init() {
	if _, err := os.Stat("public/single"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/single", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
// @Summary Upload an image
// @Description Allows a seller to upload an image for their inventory.
// @Tags sellers
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param image formData file true "Image file to be uploaded"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {string} string "Bad Request"
// @Router /sellerGroup/upload/ [post]
func (h Handler) uploadImage(c echo.Context) error {
	err := c.Request().ParseMultipartForm(50 * 1024)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))

	}
	
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))

	}

	if file.Size > 50*1024 {
		return c.String(http.StatusBadRequest, "file size exceeds 50KB")
	}

	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if fileExt != ".jpg" && fileExt != ".png" {
		return c.String(http.StatusBadRequest, "only JPG and PNG files are allowed")
	}

	fileExt = filepath.Ext(file.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:8088/images/single/" + filename

	out, err := os.Create("public/single/" + filename)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return err
	}
	defer out.Close()

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return err
	}
	defer src.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"filepath": filePath,
	})

}
// @Summary Add a new laptop for a seller
// @Description Allows a seller to add a new laptop to their inventory.
// @Tags sellers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param request body param.SellerLaptopRequest true "Laptop details to be added"
// @Success 200 {object} param.SellerLaptopRequest "Successful response"
// @Failure 400 {object} error "Bad Request"
// @Router /sellerGroup/laptops/ [post]
func (h Handler) AddSellerLaptop(c echo.Context) error {
	var req param.SellerLaptopRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	cl := claim.GetClaimFromEchoContext(c)
	req.UserID = cl.UserID

	resp, _ := h.userSvc.AddLaptop(req)

	return c.JSON(http.StatusOK, resp)
}
