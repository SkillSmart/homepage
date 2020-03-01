package logg


// Enumerate the lists of set LogLevels for the system to work with
const (
	// Define the LogLevels to be used
	LevelInfo = iota
	LevelDebug
	//TODO: Add relevant logging levels to use
)

// Define a slice of all available logLevels
var logLevels = []int{LevelInfo, LevelDebug}
