package main

import (
	"fmt"
	"net/http"

	"github.com/rokin04/go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
