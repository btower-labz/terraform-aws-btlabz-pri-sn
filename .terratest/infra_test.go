package infra_test

import (
	"github.com/btower-labz/terraform-public-modules-tests/utils"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRegionFromEnv(t *testing.T) {
	t.Parallel()
	region, ok := os.LookupEnv("TERRATEST_REGION")
	require.True(t, ok, "Variable is required: %v", "TERRATEST_REGION")
	require.NotEmpty(t, region, "Variable should not be empty: %v", "TERRATEST_REGION")
	logger.Logf(t, "Testing region: %v", region)
	utils.DoInfraDeploy(t, region, "../.infratest")
}
