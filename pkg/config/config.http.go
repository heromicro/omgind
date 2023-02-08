package config

// HTTP http配置参数
type HTTPConfig struct {
	Host             string
	Port             int
	CertFile         string
	KeyFile          string
	ShutdownTimeout  int
	MaxContentLength int64
	MaxLoggerLength  int `default:"4096"`
}

// JWTAuth 用户认证
type JWTAuthConfig struct {
	Enable        bool
	SigningMethod string
	SigningKey    string
	Expired       int
	Store         string
	FilePath      string
	RedisDB       int
	RedisPrefix   string
}

// RateLimiter 请求频率限制配置参数
type RateLimiterConfig struct {
	Enable  bool
	Count   int64
	RedisDB int
}

// CORS 跨域请求配置参数
type CORSConfig struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

// GZIP gzip压缩
type GZIPConfig struct {
	Enable             bool
	ExcludedExtentions []string
	ExcludedPaths      []string
}
