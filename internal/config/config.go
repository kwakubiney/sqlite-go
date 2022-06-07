package config
import(
	"os"
	"github.com/joho/godotenv"
)

func LoadMainConfig(path string) error {
	if _, err := os.Stat(path); err == nil {
		return godotenv.Load(path)
	}

	return nil
}