package tests

import (
	"fmt"
	"testing"

	"github.com/menxit/stairway-to-go/helloworld"
	"github.com/stretchr/testify/assert"
)

func Test_helloworld_Restituisco10ComeIntero(t *testing.T) {
	var dovreiEssereUn10Intero int = helloworld.Restituisco10ComeIntero()
	assert.Equal(t, dovreiEssereUn10Intero, 10, "Dovrebbe essere un 10")
}

func Test_helloworld_ResituiscoIlPiGreco(t *testing.T) {
	var dovreiEssereIlPiGreco float64 = helloworld.ResituiscoIlPiGreco()
	assert.Equal(t, dovreiEssereIlPiGreco, 3.14, "Dovrei essere il pi greco")
}

func Test_helloworld_RestituiscoTrue(t *testing.T) {
	var dovreiEssereVero bool = helloworld.RestituiscoTrue()
	assert.True(t, dovreiEssereVero, "Dovrei essere vero")
}

func Test_helloworld_DaFloatAIntero(t *testing.T) {
	var dovreiEssereUnIntero int = helloworld.DaFloatAIntero(3.14)
	assert.Equal(t, dovreiEssereUnIntero, 3, "Dovrei essere un intero")
}

func Test_helloworld_CercaIlMassimoInUnoSlice(t *testing.T) {
	ilMioSlice := []int{-1, 23, -3, 4, 5}
	ilMioSlice = append(ilMioSlice, 30)
	fmt.Println("Debug [CercaIlMassimoInUnoSlice] >> ", &ilMioSlice[0])
	max, _ := helloworld.CercaIlMassimoInUnoSlice(ilMioSlice)
	assert.Equal(t, max, 30, "Il massimo dovrebbe essere 30")
}

func Test_helloworld_CercaIlMassimoInUnArrayPassaggioPerValore(t *testing.T) {
	ilMioArray := [5]int{-1, 23, -3, 4, 5}
	fmt.Println("Debug [CercaIlMassimoInUnArrayPassaggioPerValore] >> ", &ilMioArray[0])
	max, err := helloworld.CercaIlMassimoInUnArrayPassaggioPerValore(ilMioArray)
	assert.Equal(t, max, 23, "Il massimo dovrebbe essere 23")
	assert.Nil(t, err)
}

func Test_helloworld_CercaIlMassimoInUnArray(t *testing.T) {
	ilMioArray := [5]int{-1, 23, -3, 4, 5}
	fmt.Println("Debug [CercaIlMassimoInUnArray] >> ", &ilMioArray[0])
	max, err := helloworld.CercaIlMassimoInUnArray(&ilMioArray)
	assert.Equal(t, max, 23, "Il massimo dovrebbe essere 23")
	assert.Nil(t, err)
}
