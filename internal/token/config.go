package token

import "time"

type Config struct {
	JWTSecret          string        `env:"JWT_SECRET"`
	JWTTokenExpiration time.Duration `env:"JWT_TOKEN_EXPIRATION" ,envDefault:"1h"`
}
