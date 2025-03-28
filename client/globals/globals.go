package globals

type Config struct { //? No se cuando se hace nill
	Ip      string `json:"ip"`
	Puerto  int    `json:"puerto"`
	Mensaje string `json:"mensaje"`
}

var ClientConfig *Config