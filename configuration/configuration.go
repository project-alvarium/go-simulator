package configuration

// For randomizer
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// [Streams Configs]
// Place announcement address here
const AnnAddress = "434e22869231fb44a9236dcd57050aa3d30e862585c44b973588f5023d0967d40000000000000000:6077dab36acb4bd8d00cf018"

// URL for author console
const AuthConsoleUrl = "http://127.0.0.1:8080"

// URL for IOTA node
const NodeUrl = "http://68.183.204.5:14265"


// Configuration holder
type Configuration struct {
	DatabaseName string
	DatabaseURL  string
	HTTPPort     string
	Secret       string
}

// Config object
var (
	Config Configuration
)

func InitConfig() {
	Config = Configuration{
		DatabaseName: "alvarium-db",
		DatabaseURL:  "mongodb://localhost:27017",
		HTTPPort:     "9091",
		Secret:       "techdev",
	}
}
