package config

// jwtConfig token 配置 接口
type JwtConfig interface {
	GetSecretKey() string
}

// defaultJwtConfig token 配置
type defaultJwtConfig struct {
	SecretKey string `json:"secretKey"`
}

// GetSecretKey token 密钥
func (m defaultJwtConfig) GetSecretKey() string {
	return m.SecretKey
}
