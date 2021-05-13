package config

type Server struct {
	Address string
}

type Store struct {
	Name string
	Dsn  string // data source name
}

type Conf struct {
	Environment string
	Server      Server
	Store       Store
}

func New() Conf {
	c := Conf{
		Environment: "prod",
		Server: Server{
			Address: ":8888",
		},
		Store: Store{
			Name: "sqlite",
			Dsn:  "./commentify.db",
		},
	}

	return c
}
