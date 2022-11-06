package config

var cfg *Config

type (
	Config struct {
		App          `yaml:"app"`
		HTTP         `yaml:"http"`
		Log          `yaml:"logger"`
		Postgres     `yaml:"postgres"`
		MongoDb      `yaml:"mongodb"`
		Security     `yaml:"security"`
		EtherealMail `yaml:"etherealMail"`
		Application  `yaml:"application"`
		Cache        `yaml:"cache"`
	}

	Cache struct {
		RedisUrl      string `env-required:"true" yaml:"url" env:"REDIS_URL"`
		RedisDb       int    `env-required:"true" env:"REDIS_DB"`
		RedisPort     int    `env-required:"true" env:"REDIS_PORT"`
		RedisPassword string `yaml:"password" `
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Application struct {
		AppWebUrl string `env-required:"true" yaml:"app_web_url" env:"APP_WEB_URL"`
	}

	Postgres struct {
		Host         string `env-required:"true" yaml:"host" env:"HOST_DB"`
		PostgresPort int    `env-required:"true" env:"PORT_DB"`
		Username     string `env-required:"true" yaml:"username" env:"USERNAME_DB"`
		Password     string `env-required:"true" yaml:"password" env:"PASSWORD_DB"`
		Dbname       string `yaml:"dbname"`
		Schema       string `yaml:"schema"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	}

	MongoDb struct {
		MongoDbDatabase string `env-required:"true" yaml:"database" env:"DATABASE_MONGODB"`
		MongoDbHost     string `env-required:"true" yaml:"host" env:"HOST_MONGODB"`
		MongoDbPort     int    `env-required:"true" env:"PORT_MONGODB"`
		MongoDbUsername string `yaml:"username" env:"USERNAME_MONGODB"`
		MongoDbPassword string `yaml:"password" env:"PASSWORD_MONGODB"`
	}

	Security struct {
		Jwt `yaml:"jwt"`
	}

	EtherealMail struct {
		Smtp `yaml:"smtp"`
	}

	Jwt struct {
		Secret string `env-required:"true" yaml:"secret" env:"JWT_SECRET"`
		Issuer string `env-required:"true" yaml:"issuer" env:"JWT_ISSUER"`
	}

	Smtp struct {
		Host     string `env-required:"true" yaml:"host" env:"HOST_ETHEREAL_MAIL"`
		SmtpPort string `env-required:"true" yaml:"port" env:"PORT_ETHEREAL_MAIL"`
		Username string `env-required:"true" yaml:"username" env:"USERNAME_ETHEREAL_MAIL"`
		Password string `env-required:"true" yaml:"password" env:"PASSWORD_ETHEREAL_MAIL"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
