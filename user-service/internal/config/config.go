package config

// Захардкодил для простоты примера
const (
	port        = 8080
	postgresDsn = "host=localhost port=5432 user=user password=password dbname=postgres sslmode=disable"
	billingDSN  = "http://localhost:8085"
)

type Config struct {
	Port        int
	PostgresDSN string
	BillingDSN  string
}

func New() Config {
	return Config{
		Port:        port,
		PostgresDSN: postgresDsn,
		BillingDSN:  billingDSN,
	}
}
