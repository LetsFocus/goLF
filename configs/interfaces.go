package configs

type configs interface {
	Get(key string) string
	GetOrDefault(key, defaultValue string) string
}
