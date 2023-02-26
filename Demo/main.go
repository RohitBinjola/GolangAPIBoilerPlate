// How to implement Boiler Plate structure for an api in GIN in real world.

package main

import (
	"Demo/src/config"
	"Demo/src/logger"
	"Demo/src/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")
	// database.Init()
	// if config.Appconfig.GetBool("seeddata") {
	// 	Logic to seed data to database
	// 	logger.InfoLn("Data seeded successfully")
	// }
	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
