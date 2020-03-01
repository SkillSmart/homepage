
// This setup allows to use chained access to predefined, leveled logging messages along system domains
//
// Examples:::
// logg.Config.SetKeyInfo()
// logg.Config.SetKeyFatal()


package logg

var (
	Config loggerInterface
	Application loggerInterface
	Environment loggerInterface
)

func init() {
	// Initiate the default loggers for the system
	Config = configLogger
	Application = applicationLogger
	Environment = environmentLogger
}

// This defines the public interface that is accessible for the internal logging package
type loggerInterface interface {
	Info(msg string)
	Fatal(err error)
}

type logResponse struct {
	Message string `json:"message"`
	// TODO: Define additional data
}


