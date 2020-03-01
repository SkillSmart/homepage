package logg

import "strconv"


// IsLogLevel checks if a given value represents a valid LogLevel
func IsLogLevel(level interface{}) bool {
	switch v := level.(type) {
	case string:
		var err error
		// We first need to convert to int
		level, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return false
		}
	default:
		return false
	}

	// Now check if the level is contained in the given list of values
	for _, lvl := range logLevels {
		if lvl == level {
			return true
		}
	}
	return false
}

