package util_test

import (
	"testing"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-common/util"
	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {

	cases := [][]int{
		{1, 5, 8, 2, 6, 7, 3, 4},
		{1, 5, 8, 2, 6, 3, 4},
	}

	for i := 0; i < len(cases); i++ {
		t.Log("----", cases[i])
		r := util.RangeSet{}
		for _, v := range cases[i] {
			err := r.Add(v, util.Consecutive, true)
			require.NoError(t, err)
			t.Log(r)
		}

		t.Log(r.Defragment())
	}

}
