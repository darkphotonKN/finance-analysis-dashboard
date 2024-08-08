package config

import "os"

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
