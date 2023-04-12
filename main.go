package main

import (
	"firstDevProject/cron"
	initializers "firstDevProject/initializeres" // this is the way to name the imports
	"firstDevProject/routes"
)

func init() { //runs before main()
	initializers.LoadEnvVariables()
	cron.StartTestCronJob()
	// connections.ConnectToDB()
}

func main() {
	r := routes.SuperRouter()
	r.Run()
}
