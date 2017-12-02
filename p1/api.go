package p1

import (
	"net/http"

	"github.com/labstack/echo"
)

func Pay(c echo.Context) error {
	// f := new(Fruit)
	// if err := c.Bind(f); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"errInfo": err.Error()})
	// }
	fruit := Fruit{}
	fruit.Name = "li"
	fruit.Prict = 20
	return c.JSON(http.StatusOK, fruit)
}
