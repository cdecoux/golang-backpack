package pkg

import "time"

func TimeOf(f func() interface{}) (time.Duration, interface{}) {
	start := time.Now()
	returnValue := f()
	end := time.Now()
	return end.Sub(start), returnValue
}

// Seriously, GoLang needs overloading capabilities
func TimeOfVoid(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Now()
	return end.Sub(start)
}