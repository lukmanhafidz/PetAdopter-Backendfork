package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Appconfig struct {
	Username   string
	Password   string
	Address    string
	Port       int
	DBName     string
	fPath      string
	bucketName string
	projectID  string
}

var (
	lock      = &sync.Mutex{}
	appconfig *Appconfig
)

func Getconfig() *Appconfig {
	lock.Lock()
	defer lock.Unlock()

	if appconfig == nil {
		appconfig = initConfig()
	}

	return appconfig
}

func initConfig() *Appconfig {
	var defaultconfig Appconfig

	// err := godotenv.Load("local.env")
	// if err != nil {
	// 	log.Println("cant load env file")
	// 	return nil
	// }

	defaultconfig.Username = os.Getenv("Username")
	defaultconfig.Password = os.Getenv("Password")
	defaultconfig.Address = os.Getenv("Address")

	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Println("Cant convert port")
		return nil
	}

	defaultconfig.Port = port
	defaultconfig.fPath = os.Getenv("fPath")
	defaultconfig.bucketName = os.Getenv("bucketName")
	defaultconfig.projectID = os.Getenv("projectID")
	defaultconfig.DBName = os.Getenv("DBName")

	return &defaultconfig
}
