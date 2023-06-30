package storage

type Config struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	DbName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}
