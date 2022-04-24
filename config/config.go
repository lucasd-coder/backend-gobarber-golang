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
		RedisUrl      string `yaml:"url"`
		RedisDb       int    `yaml:"db"`
		RedisPassword string `yaml:"password"`
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
		Host         string `env-required:"true" yaml:"host"`
		PostgresPort int    `env-required:"true" yaml:"port"`
		Username     string `env-required:"true" yaml:"username" env:"USERNAME_DB"`
		Password     string `env-required:"true" yaml:"password" env:"PASSWORD_DB"`
		Dbname       string `env-required:"true" yaml:"dbname"`
		Schema       string `env-required:"true" yaml:"schema"`
		MaxIdleConns int    `env-required:"true" yaml:"maxIdleConns"`
		MaxOpenConns int    `env-required:"true" yaml:"MaxOpenConns"`
	}

	MongoDb struct {
		MongoDbHost string `env-required:"true" yaml:"host"`
		MongoDbPort int    `env-required:"true" yaml:"port"`
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
