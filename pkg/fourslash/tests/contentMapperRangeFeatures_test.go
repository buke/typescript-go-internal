package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/testutil"
	"github.com/buke/typescript-go-internal/pkg/testutil/contentmappertest"
)

func TestContentMapperRangeFeaturesIncludeScriptInsideMarkup(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := newContentMapperFourslash(t, `// @Filename: /app.vue
<template>before</template>
<script lang="ts">
const message = "world";
message/*selection*/;
</script>
<template>after</template>
`, contentmappertest.ComponentMapper, ".vue")
	defer done()

	// Selection ranges at the marker must be built from mappable virtual ancestors even though the virtual
	// file contains a synthesized render-function suffix.
	f.VerifyBaselineSelectionRanges(t)
}
