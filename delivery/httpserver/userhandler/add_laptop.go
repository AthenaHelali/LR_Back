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

func (h Handler) uploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))

	}

	fileExt := filepath.Ext(file.Filename)
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
