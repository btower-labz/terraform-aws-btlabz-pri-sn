package test

import (
	"os"
	"path/filepath"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/logger"

	"github.com/stretchr/testify/require"
	"testing"
)

const TERRATEST_REGION_KEY = "TERRATEST_REGION"

func GetAWSRegionFromEnvironment(t *testing.T) string {
	region, ok := os.LookupEnv(TERRATEST_REGION_KEY)
	require.True(t, ok, "Variable is required: %v", TERRATEST_REGION_KEY)
	require.NotEmpty(t, region, "Variable should not be empty: %v", TERRATEST_REGION_KEY)
	logger.Logf(t, "Testing region: %v", region)
	return region
}

// Subdirectiry under the main module
const INFRA_TEST_DIR_PART = ".infratest"

// Copy module root
// Add tests to .infratest
// Return .infratest path
func CopyInfraTestFolderToTemp(t *testing.T, rootModuleDir string, infraTestDir string) string {

	require.NotEmpty(t, infraTestDir)
	require.DirExists(t, infraTestDir)

	require.NotEmpty(t, rootModuleDir)
	require.DirExists(t, rootModuleDir)

	// The function with the similar name in test_structure has SKIP_* side effect.
	tempModuleRoot, err := files.CopyTerraformFolderToTemp(rootModuleDir, "terratest")
	require.NoError(t, err)
	require.NotEmpty(t, tempModuleRoot)
	require.DirExists(t, tempModuleRoot)

	logger.Logf(t, "Copied module root from %s to %s", rootModuleDir, tempModuleRoot)

	err = files.CopyFolderContents(infraTestDir, filepath.Join(tempModuleRoot, INFRA_TEST_DIR_PART))
	require.NoError(t, err)
	tempInfraTest := filepath.Join(tempModuleRoot, INFRA_TEST_DIR_PART)
	require.NotEmpty(t, tempInfraTest)
	require.DirExists(t, tempInfraTest)
	logger.Logf(t, "Copied infratest root from %s to %s", infraTestDir, tempInfraTest)

	return tempInfraTest
}
