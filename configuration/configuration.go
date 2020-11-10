package configuration

// For randomizer
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// [Streams Configs]
// Place announcement address here
const AnnAddress = "caa066b039d1fd0fc35327aab4d8ed750f015ee3f88254044859d350595db68f0000000000000000:ace47289e46d12256ef9b368"
// URL for author console
const AuthConsoleUrl = "http://127.0.0.1:8080"
// URL for IOTA node
const NodeUrl = "http://localhost:14265"
// Min Weight Magnitude
const NodeMwm = 9


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
		HTTPPort:     "9090",
		Secret:       "techdev",
	}
}
