package config

type (
	Config struct {
		App          App            `yaml:"app"`
		Postgres     Postgres       `yaml:"postgres"`
		Account      Account        `yaml:"account"`
		DataProtocol DataProtocol   `yaml:"data"`
		Auth         Authentication `yaml:"jwt"`
	}

	App struct {
		Name    string `yaml:"name" envconfig:"APP_NAME"`
		Address string `yaml:"port" envconfig:"APP_ADDRESS"`
	}

	Postgres struct {
		Username string `yaml:"username" envconfig:"POSTGRES_USERNAME"`
		Password string `yaml:"password" envconfig:"POSTGRES_PASSWORD"`
		DBName   string `yaml:"db_name" envconfig:"POSTGRES_DBNAME"`
		Host     string `yaml:"host" envconfig:"POSTGRES_HOST"`
		Port     string `yaml:"port" envconfig:"POSTGRES_PORT"`
	}

	Authentication struct {
		Secret            string `yaml:"secret"`
		TokenExpireTime   int    `yaml:"token_expire_time"`
		RefreshExpireTime int    `yaml:"refresh_expire_time"`
	}

	Account struct {
		MinUsernameLength int `yaml:"min_username_length"`
	}

	DataProtocol struct {
		PageSize uint `yaml:"page_size"`
	}
)
