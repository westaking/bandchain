package bandrng_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/bandchain/chain/pkg/bandrng"
)

func TestChooseOneOne(t *testing.T) {
	r := bandrng.NewRng("SEED")
	weights := []uint64{10, 13, 10, 25, 42} // prefix sum is 10,23,33,58,100

	require.Equal(t, bandrng.ChooseOne(r, weights), 4) // rng NextUint64() will return 15735084640102210068
	require.Equal(t, bandrng.ChooseOne(r, weights), 4) // rng NextUint64() will return 3485776390957061973
	require.Equal(t, bandrng.ChooseOne(r, weights), 3) // rng NextUint64() will return 17609118114147816341
	require.Equal(t, bandrng.ChooseOne(r, weights), 2) // rng NextUint64() will return 15960811988050104523
	require.Equal(t, bandrng.ChooseOne(r, weights), 3) // rng NextUint64() will return 11919533627209787235

	r = bandrng.NewRng("SEED")
	weights = []uint64{2, 4, 4} // prefix sum is 2,6,10

	require.Equal(t, bandrng.ChooseOne(r, weights), 2) // rng NextUint64() will return 15735084640102210068
	require.Equal(t, bandrng.ChooseOne(r, weights), 1) // rng NextUint64() will return 3485776390957061973
	require.Equal(t, bandrng.ChooseOne(r, weights), 0) // rng NextUint64() will return 17609118114147816341

}

func TestChooseOnePanic(t *testing.T) {
	r := bandrng.NewRng("SEED")
	require.Panics(t, func() {
		bandrng.ChooseOne(r, []uint64{math.MaxUint64, math.MaxUint64})
	})
	require.Panics(t, func() {
		bandrng.ChooseOne(r, []uint64{1, math.MaxUint64})
	})

	require.Panics(t, func() {
		bandrng.ChooseOne(r, []uint64{math.MaxUint64, 1})
	})
}

func TestChooseSomeEqualWeights(t *testing.T) {
	r := bandrng.NewRng("SEED")
	length := 10
	weights := make([]uint64, length)
	for idx := 0; idx < length; idx++ {
		weights[idx] = 1
	}
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{8, 0, 6, 7, 4})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{2, 7, 4, 8, 9})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{6, 0, 9, 5, 3})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{2, 7, 0, 3, 9})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{8, 3, 4, 0, 1})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{6, 7, 0, 4, 9})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{0, 4, 5, 2, 1})
}

func TestChooseSomeSkewedWeights(t *testing.T) {
	r := bandrng.NewRng("SEED")
	length := 10
	weights := make([]uint64, length)
	for idx := 0; idx < length; idx++ {
		weights[idx] = uint64(1 + idx*10)
	}
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{7, 4, 9, 6, 3})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{6, 9, 5, 8, 3})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{2, 6, 8, 1, 3})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{8, 2, 9, 7, 4})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{8, 9, 7, 6, 4})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{7, 5, 8, 9, 6})
	require.Equal(t, bandrng.ChooseSome(r, weights, 5), []int{5, 8, 6, 7, 9})
}

func TestChooseSomeMaxWeight(t *testing.T) {
	r := bandrng.NewRng("SEED")
	length := 10
	weights := make([]uint64, length)
	for idx := 0; idx < length; idx++ {
		weights[idx] = uint64(1 + idx*10)
	}
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{6, 9, 5, 8, 3})
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{7, 5, 8, 9, 6})
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{5, 8, 6, 7, 9})
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{6, 8, 9, 7, 4})
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{7, 9, 3, 8, 2})
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{1, 7, 8, 9, 5})
	require.Equal(t, bandrng.ChooseSomeMaxWeight(r, weights, 5, 3), []int{9, 8, 6, 5, 3})
}
