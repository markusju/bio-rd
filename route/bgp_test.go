package route

import (
	"testing"

	"github.com/bio-routing/bio-rd/protocols/bgp/types"
	"github.com/stretchr/testify/assert"
)

func TestComputeHash(t *testing.T) {
	p := &BGPPath{
		ASPath: types.ASPath{
			types.ASPathSegment{
				ASNs: []uint32{123, 456},
				Type: types.ASSequence,
			},
		},
		BGPIdentifier: 1,
		Communities: []uint32{
			123, 456,
		},
		EBGP: false,
		LargeCommunities: []types.LargeCommunity{
			types.LargeCommunity{
				DataPart1:           1,
				DataPart2:           2,
				GlobalAdministrator: 3,
			},
		},
		LocalPref:      100,
		MED:            1,
		NextHop:        100,
		PathIdentifier: 5,
		Source:         4,
	}

	assert.Equal(t, "45e238420552b88043edb8cb402034466b08d53b49f8e0fedc680747014ddeff", p.ComputeHash())

	p.LocalPref = 150

	assert.NotEqual(t, "45e238420552b88043edb8cb402034466b08d53b49f8e0fedc680747014ddeff", p.ComputeHash())
}
