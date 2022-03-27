package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string
	DBSource      string
	ServerAddress string
}

func LoadConfig(path string) (config Config, err error) {
	vi := viper.New()
	vi.AddConfigPath(path)
	vi.SetConfigName("config")
	vi.SetConfigType("yaml")

	vi.AutomaticEnv()
	err = vi.ReadInConfig()
	if err != nil {
		return
	}
	config = Config{
		DBDriver:      vi.GetString("DB_DRIVER"),
		DBSource:      vi.GetString("DB_SOURCE"),
		ServerAddress: vi.GetString("SERVER_ADDRESS"),
	}
	return
}
