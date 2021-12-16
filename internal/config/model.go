package config

type (
	Config struct {
		App      App `json:"app"`
		Database struct {
			Postgres Postgres `json:"postgres"`
		} `json:"database"`
		DataProtocol   DataProtocol   `json:"data_protocol"`
		Authentication Authentication `json:"authentication"`
		Authorization  Authorization  `json:"authorization"`
	}

	ConfigurationType struct {
		Dev    Config `json:"dev"`
		Prod   Config `json:"prod"`
		IsProd bool   `json:"is_prod"`
	}

	Authorization struct {
		Casbin struct {
			ModelPath  string `json:"model_path"`
			PolicyPath string `json:"policy_path"`
		} `json:"casbin"`
	}

	App struct {
		Name string `json:"name"`
		Port string `json:"port"`
	}

	Postgres struct {
		Username      string `json:"username"`
		Password      string `json:"password"`
		DBName        string `json:"db_name"`
		Host          string `json:"host"`
		Port          string `json:"port"`
		AutoMigration bool   `json:"auto_migration"`
	}

	Authentication struct {
		Secret            string `json:"secret"`
		TokenExpireTime   int    `json:"token_expire_time"`
		RefreshExpireTime int    `json:"refresh_expire_time"`
	}

	DataProtocol struct {
		PageSize uint `json:"page_size"`
	}
)
