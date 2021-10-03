package math

import (
	log "github.com/sirupsen/logrus"
	"math/big"
)

func FactorialParallel(factorial int, parallelism int) *big.Int {

	result := big.NewInt(1)

	factChan := make(chan *big.Int, parallelism)


	for i := 1; i <= parallelism; i++ {
		go func(start int) {
			partialFactorial := big.NewInt(1)

			for i := start; i <= factorial; i += parallelism {
				partialFactorial.Mul(partialFactorial, big.NewInt(int64(i)))
			}
			factChan <- partialFactorial
		}(i)
	}

	for i := 0; i < parallelism; i++ {
		partialFactorial := <-factChan
		log.Debug("Received partial factorial: ", partialFactorial)
		result.Mul(result, partialFactorial)
	}

	log.Debug("Final Result: ", result)

	return result
}