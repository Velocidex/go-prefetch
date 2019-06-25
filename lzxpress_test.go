package prefetch

import (
	"encoding/json"
	"os"
	"path"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

var test_cases = []string{
	"test_data/Prefetch/Prefetch.Test/TestFiles/Win10/CALC.EXE-3FBEF7FD.pf",
	"test_data/Prefetch/Prefetch.Test/TestFiles/Win8x/LIVECOMM.EXE-D546E475.pf",
	"test_data/Prefetch/Prefetch.Test/TestFiles/Win7/PING.EXE-B29F6629.pf",
	"test_data/Prefetch/Prefetch.Test/TestFiles/XPPro/MSMSGS.EXE-2B6052DE.pf",
}

func TestPrefetch(t *testing.T) {
	for _, test_case := range test_cases {
		fd, err := os.Open(test_case)
		assert.NoError(t, err)

		prefetch_info, err := LoadPrefetch(fd)
		assert.NoError(t, err)

		serialized, err := json.MarshalIndent(prefetch_info, " ", " ")
		assert.NoError(t, err)

		goldie.Assert(t, path.Base(test_case), serialized)
	}
}

func init() {
	time.Local = time.UTC
	spew.Config.DisablePointerAddresses = true
	spew.Config.SortKeys = true
}
