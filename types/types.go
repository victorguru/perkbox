package types

// Environment holds the environment variables coming from the docker-compose file
type Environment struct {
	DSNDB         string
	MigrationsDir string
}

// Coupon is the structure for a coupon
type Coupon struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Brand     string  `json:"brand"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"created_at"`
	Expiry    string  `json:"expiry"`
}
