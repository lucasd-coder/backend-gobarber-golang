package config

var cfg *Config

type (
	Config struct {
		App          `yaml:"app"`
		HTTP         `yaml:"http"`
		Log          `yaml:"logger"`
		Postgres     `yaml:"postgres"`
		Security     `yaml:"security"`
		EtherealMail `yaml:"etherealMail"`
		Application  `yaml:"application"`
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
		Username     string `env-required:"true" yaml:"username"`
		Password     string `env-required:"true" yaml:"password"`
		Dbname       string `env-required:"true" yaml:"dbname"`
		Schema       string `env-required:"true" yaml:"schema"`
		MaxIdleConns int    `env-required:"true" yaml:"maxIdleConns"`
		MaxOpenConns int    `env-required:"true" yaml:"MaxOpenConns"`
	}

	Security struct {
		Jwt `yaml:"jwt"`
	}

	EtherealMail struct {
		Smtp `yaml:"smtp"`
	}

	Jwt struct {
		Secret string `env-required:"true" yaml:"secret"`
		Issuer string `env-required:"true" yaml:"issuer"`
	}

	Smtp struct {
		Host     string `env-required:"true" yaml:"host"`
		SmtpPort string `env-required:"true" yaml:"port"`
		Username string `env-required:"true" yaml:"username"`
		Password string `env-required:"true" yaml:"password"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
