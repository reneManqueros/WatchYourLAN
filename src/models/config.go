package models

import "github.com/spf13/viper"

type Conf struct {
	DbProvider           string
	Iface                string
	DbPath               string
	GuiIP                string
	GuiPort              string
	GuiAuth              string
	ShoutUrl             string
	Theme                string
	MongoDBConnectionURI string
	MongoDBDatabase      string
	MongoDBCollection    string
	Timeout              int
}

var AppConfig Conf

const configPath = "/data/config"

func (c *Conf) Get() {
	viper.SetDefault("IFACE", "enp1s0")
	viper.SetDefault("DBPATH", "/data/db.sqlite")
	viper.SetDefault("DBPROVIDER", "sqlite")
	viper.SetDefault("GUIIP", "localhost")
	viper.SetDefault("GUIPORT", "8840")
	viper.SetDefault("GUIAUTH", "")
	viper.SetDefault("TIMEOUT", "60")
	viper.SetDefault("SHOUTRRR_URL", "")
	viper.SetDefault("THEME", "solar")
	viper.SetDefault("MONGODBCONNECTIONURI", "")
	viper.SetDefault("MONGODBDATABASE", "")
	viper.SetDefault("MONGODBCOLLECTION", "")

	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	c.Iface = viper.Get("IFACE").(string)
	c.DbPath = viper.Get("DBPATH").(string)
	c.GuiIP = viper.Get("GUIIP").(string)
	c.GuiPort = viper.Get("GUIPORT").(string)
	c.GuiAuth = viper.Get("GUIAUTH").(string)
	c.Timeout = viper.GetInt("TIMEOUT")
	c.ShoutUrl = viper.Get("SHOUTRRR_URL").(string)
	c.Theme = viper.Get("THEME").(string)
	c.DbProvider = viper.Get("DBPROVIDER").(string)
	c.MongoDBDatabase = viper.Get("MONGODBDATABASE").(string)
	c.MongoDBConnectionURI = viper.Get("MONGODBCONNECTIONURI").(string)
	c.MongoDBCollection = viper.Get("MONGODBCOLLECTION").(string)
}

func (c *Conf) Set() {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.Set("THEME", AppConfig.Theme)
	viper.WriteConfig()
}
