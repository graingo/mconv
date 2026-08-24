package mconv_test

import (
	"testing"

	"github.com/graingo/mconv"
)

// TestDeprecatedV1CacheFunctionsRemainCallable makes the v1 compatibility
// promise explicit. Their removal belongs to the /v2 module described in
// V2_MIGRATION.md and must be rejected by API compatibility CI on v1.
func TestDeprecatedV1CacheFunctionsRemainCallable(_ *testing.T) {
	mconv.SetTypeInfoCacheSize(1)
	mconv.SetConversionCacheSize(1)
	mconv.ClearTypeInfoCache()
	mconv.ClearConversionCache()
}
