package configuration

// For randomizer
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// [Streams Configs]
// Place announcement address here
const AnnAddress = "ARK9ZOGNCWEONTMOYRYYNLLG9JPGBSTFVCHSFIKQFS9XFYKQDMSFPTGXUGUSHLZ9VZXAOBFTCKHVJRAFW:2779530283277761"

// URL for author console
const AuthConsoleUrl = "http://127.0.0.1:8080"

// URL for IOTA node
const NodeUrl = "http://localhost:15601"

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
