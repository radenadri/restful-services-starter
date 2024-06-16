package config

import "boilerplate/app/utils"

var DB_HOST = utils.LoadEnv("DB_HOST")
var DB_PORT = utils.LoadEnv("DB_PORT")
var DB_NAME = utils.LoadEnv("DB_NAME")
var DB_USER = utils.LoadEnv("DB_USER")
var DB_PASSWORD = utils.LoadEnv("DB_PASSWORD")
var DB_SSL_MODE = utils.LoadEnv("DB_SSL_MODE")
