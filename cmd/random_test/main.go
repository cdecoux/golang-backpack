package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	sandboxTools "sandbox/pkg"
	mathTools "sandbox/pkg/math"
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
	// Get a Map of Weights
	mapOfWeights := make(map[string]int)
	mapOfWeights["a"] = 10 // 25%
	mapOfWeights["b"] = 20 // 50%
	mapOfWeights["c"] = 10 // 25%

	ds, err := mathTools.NewDistributionSelector(mapOfWeights)
	if err != nil {
		return
	}

	singleThreadedTime, singleThreadedResults :=
		sandboxTools.TimeOf(func() interface{} {
			// Init a Map of Counts
			mapOfCounts := make(map[string]int)

			for i := 0; i < 10000000; i++ {
				selection, err := ds.SelectRandom()
				if err != nil {
					return err
				}
				mapOfCounts[selection]++
			}

			return mapOfCounts
	})

	multiThreadedTime, multiThreadedResults :=
		sandboxTools.TimeOf(func() interface{} {
			// Init a Map of Counts
			mapOfCounts := make(map[string]int)
			results, err := ds.SelectRandomN(10000000, 5)
			if err != nil {
				log.Error(err)
				return err
			}

			for _, result := range results {
				mapOfCounts[result]++
			}

			return mapOfCounts
		})

	multiThreaded1Time, multiThreaded1Results :=
		sandboxTools.TimeOf(func() interface{} {
			// Init a Map of Counts
			mapOfCounts := make(map[string]int)
			results, err := ds.SelectRandomN(10000000, 1)
			if err != nil {
				log.Error(err)
				return err
			}

			for _, result := range results {
				mapOfCounts[result]++
			}

			return mapOfCounts
		})


	log.Info("Single-Threaded Results: ", singleThreadedTime, " : ", singleThreadedResults)
	log.Info("Multi-Threaded (5) Results: ", multiThreadedTime, " : ", multiThreadedResults)
	log.Info("Multi-Threaded (1) Results: ", multiThreaded1Time, " : ", multiThreaded1Results)


}
