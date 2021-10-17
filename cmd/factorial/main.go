package main

import (
	backpack "backpack/pkg"
	backpackMath "backpack/pkg/math"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {

	/*
		Set the Log Level of LogRUS
	 */
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "INFO"
	}

	logLevel, err := log.ParseLevel(lvl)
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)

	// the first argument is always program name
	programName := os.Args[0]
	log.Debug(programName)

	factorial, _ := strconv.Atoi(os.Args[1])
	parallelism, _ := strconv.Atoi(os.Args[2])

	factorialTime, factorialResult := backpack.TimeOf(func() interface{} {
		return backpackMath.FactorialParallel(factorial, parallelism)
	})

	log.Debug(factorialResult)
	log.Info(factorialTime)
}