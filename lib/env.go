package lib

import "os"

var (
	Port string = os.Getenv("PORT")
	Mong string = os.Getenv("MONGODB_URI")
)
