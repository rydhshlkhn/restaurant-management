package main

import (
	"github.com/rydhshlkhn/restaurant-management/codebase/app"
	restaurantmanagement "github.com/rydhshlkhn/restaurant-management/internal"
)

func main() {
	app.New(restaurantmanagement.NewService()).Run()
}
