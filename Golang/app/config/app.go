package config

import "boilerplate/app/utils"

var APP_HOST = utils.LoadEnv("APP_HOST")
var APP_PORT = utils.LoadEnv("APP_PORT")

var CORS_ALLOWED_ORIGINS = utils.LoadEnv("CORS_ALLOWED_ORIGINS")

var JWT_SECRET = utils.LoadEnv("JWT_SECRET")
