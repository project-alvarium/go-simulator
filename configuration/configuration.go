package configuration

// For randomizer
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// [Streams Configs]
// Place announcement address here
const AnnAddress = "80afefb17b06ebc8379eb49679f15cc8afad74272be8d73bfbe7eadaa35959b70000000000000000:8feb38922ab617a43cc2064f"

// URL for author console
const AuthConsoleUrl = "http://127.0.0.1:8080"

// URL for IOTA node
const NodeUrl = "http://localhost:14265"

// Min Weight Magnitude
const NodeMwm = 9

// Max number of readings to conduct
const MaxReadings = 100

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
