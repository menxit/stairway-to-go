package helloworld

import (
	"errors"
	"fmt"
)

// Restituisco10ComeIntero restituisce un 10 intero
func Restituisco10ComeIntero() int {
	var sonoUnDieciIntero int = 10
	return sonoUnDieciIntero
}

// ResituiscoIlPiGreco restituisce il pigreco
func ResituiscoIlPiGreco() float64 {
	sonoIlPiGreco := 3.14
	return sonoIlPiGreco
}

// RestituiscoTrue restituisce true
func RestituiscoTrue() bool {
	var sonoVero bool = true
	return sonoVero
}

// DaFloatAIntero prende un float64 e lo trasforma in intero
func DaFloatAIntero(sonoUnFloat float64) int {
	return int(sonoUnFloat)
}

// CercaIlMassimoInUnoSlice prende un array e restituisce il valore massimo
func CercaIlMassimoInUnoSlice(sliceDiInteri []int) (int, error) {
	if len(sliceDiInteri) == 0 {
		return 0, errors.New("Dammi un array con almeno un valore")
	}
	var max int = sliceDiInteri[0]
	for i := 0; i < len(sliceDiInteri); i++ {
		if sliceDiInteri[i] > max {
			max = sliceDiInteri[i]
		}
	}
	fmt.Println("Debug [CercaIlMassimoInUnoSlice] >> ", &sliceDiInteri[0])
	return max, nil
}

// CercaIlMassimoInUnArrayPassaggioPerValore prende un array e restituisce il valore massimo
func CercaIlMassimoInUnArrayPassaggioPerValore(arrayDiInteri [5]int) (int, error) {
	// casta array in slice
	return CercaIlMassimoInUnoSlice(arrayDiInteri[:])
}

// CercaIlMassimoInUnArray prende un array e restituisce il valore massimo
func CercaIlMassimoInUnArray(arrayDiInteri *[5]int) (int, error) {
	return CercaIlMassimoInUnoSlice((*arrayDiInteri)[:])
}
