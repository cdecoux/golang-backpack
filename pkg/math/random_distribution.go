package math

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

type distributionSelector struct {
	distribution map[interface{}]int
	cachedSum int
}

/*
	Creates a distribution selector based off of a map of weights.
	Distribution values must all be positive integers. Some may be 0, but not all of them can be.
	Should there be an issue with the values, an error will be returned.
	An optional integer can be passed in to set the parallelism (used in selectRandomN)
*/
func NewDistributionSelector(distribution map[interface{}]int) (*distributionSelector, error) {
	ds := &distributionSelector{
		distribution: distribution,
		cachedSum: 0,
	}

	// Calculate sum (which is cached). Return error should this error
	_, err := ds.getSum()
	if err != nil {
		return nil, err
	}
	// Else, return distribution selector
	return ds, nil
}

func (self *distributionSelector) getSum() (int, error)  {

	if self.cachedSum == 0 { // If cachedSum hasn't been calculated yet
		for _, weight := range self.distribution {

			// If a negative number was found
			if weight < 0 {
				self.cachedSum = -1
				return -1, errors.New("distribution weights are not all positive integers")
			}
			// Add weight to sum
			self.cachedSum += weight
		}
	} else if self.cachedSum < 0 { // If cached sum was previously calculated and wasn't positive
		return -1, errors.New("distribution weights are not a positive integer")
	}

	return self.cachedSum, nil
}

func (self *distributionSelector) SelectRandom() (interface{}, error) {

	sum, err := self.getSum()
	if err != nil {
		return nil, err
	}

	selection := rand.Intn(sum) + 1
	limit := 0
	for key, weight := range self.distribution {
		limit += weight
		if selection <= limit {return key, nil}
	}

	log.Error("Was not able to get a selection!")
	log.Errorf("Distribution: %s", self.distribution)
	log.Errorf("Random selection: %d", selection)

	panic("Was not able to get a selection!")
}


/*
	Get multiple random selections. 'num' specifies how many.
	Apparently this is slower the more parallelism you put in. (not just due to limited cores)
*/
func (self *distributionSelector) SelectRandomN(num int, parallelism int) ([]interface{}, error) {

	// Init slice for holding resulting random selections
	results := make([]interface{}, 0, num)
	resultsPartialChan := make(chan []interface{}, parallelism)
	errorChan := make(chan error)

	for i := 1; i <= parallelism; i++ {
		go func(start int) {
			resultsPartial := make([]interface{}, 0, (num/parallelism) + 1)

			for i := start; i <= num; i += parallelism {
				selection, err := self.SelectRandom()
				if err != nil {
					errorChan <- err
					return
				}

				resultsPartial = append(resultsPartial, selection)
			}

			resultsPartialChan <- resultsPartial

		}(i)
	}

	for len(results) < num {
		select {
		case resultPartial := <- resultsPartialChan:
			results = append(results, resultPartial...)
		case err := <- errorChan:
			return nil, err
		}
	}

	return results, nil
}
