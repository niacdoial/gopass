package pwrules

import (
	"context"
	"testing"

	"github.com/gopasspw/gopass/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadCustomRules(t *testing.T) {
	t.Parallel()

	cfg := config.NewNoWrites()
	aliases := map[string]string{
		"real.com": "alias.com",
		"real.de":  "copy.de",
	}

	for k, v := range aliases {
		assert.NoError(t, cfg.Set("", "domain-alias."+k+".insteadOf", v))
	}

	ctx := context.Background()
	ctx = cfg.WithConfig(ctx)

	a := LookupAliases(ctx, "alias.com")
	assert.Equal(t, []string{"real.com"}, a)

	a = LookupAliases(ctx, "copy.de")
	assert.Equal(t, []string{"real.de"}, a)

	assert.Greater(t, len(AllAliases(ctx)), 256)
}
