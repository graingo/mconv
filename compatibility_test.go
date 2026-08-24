package mconv_test

import (
	"testing"

	"github.com/graingo/mconv"
)

// TestDeprecatedV1CacheFunctionsRemainCallable makes the v1 source-compatibility
// promise explicit. API compatibility CI must reject their removal from v1.
func TestDeprecatedV1CacheFunctionsRemainCallable(_ *testing.T) {
	mconv.SetTypeInfoCacheSize(1)
	mconv.SetConversionCacheSize(1)
	mconv.ClearTypeInfoCache()
	mconv.ClearConversionCache()
}
