package test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/stretchr/testify/require"
	"testing"
)

// Deploy and test infra in the
func TestRegionFromEnv(t *testing.T) {
	t.Parallel()

	// Test case directory
	testCase := "default"

	// Get region from env
	region := GetAWSRegionFromEnvironment(t)

	// Make a copy of the root module
	_infraTestDir := CopyInfraTestFolderToTemp(t, "../", "../.infratest")
	defer os.RemoveAll(_infraTestDir)

	// Make sure the test case is in place
	infraTestDir := filepath.Join(_infraTestDir, testCase)
	require.NotEmpty(t, infraTestDir)
	require.DirExists(t, infraTestDir)
	logger.Logf(t, "Terraform test case directory: %v", infraTestDir)

	// Configure variablese
	test_structure.RunTestStage(t, "configure", func() {

		// Generate unique run identifier
		uniqueId := random.UniqueId()
		logger.Logf(t, "uniqueId: %v", uniqueId)

		// Setup common environment for all the states
		infraEnvVars := map[string]string{
			"AWS_REGION":              region,
			"TF_INPUT":                "0",
			"TF_VAR_terratest_run_id": uniqueId,
			"TF_VAR_region":           region,
			"TF_IN_AUTOMATION":        "YES",
			"TF_CLI_ARGS_plan":        "-compact-warnings",
			"TF_CLI_ARGS_apply":       "-compact-warnings",
		}

		// Terraform module variables
		tfVars := map[string]interface{}{
			"name_prefix": fmt.Sprintf("terratest-%s", uniqueId),
		}

		// See: https://github.com/gruntwork-io/terratest/blob/master/modules/terraform/options.go
		terraformOptions := &terraform.Options{
			NoColor:      true,
			Vars:         tfVars,
			TerraformDir: infraTestDir,
			EnvVars:      infraEnvVars,
		}
		test_structure.SaveTerraformOptions(t, infraTestDir, terraformOptions)
	})

	// Do init\plan
	test_structure.RunTestStage(t, "plan", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, infraTestDir)
		terraform.Init(t, terraformOptions)
		terraform.Plan(t, terraformOptions)
	})

	// Do destroy
	defer test_structure.RunTestStage(t, "destroy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, infraTestDir)
		terraform.Destroy(t, terraformOptions)
	})

	// Deploy infrastructure
	test_structure.RunTestStage(t, "apply", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, infraTestDir)
		// terraform.ApplyAndIdempotent(t, terraformOptions)
                terraform.Apply(t, terraformOptions)
	})

}
