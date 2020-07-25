package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func EnvWithNewVar(env map[string]string, name string, value string) map[string]string {
	newEnv := make(map[string]string)
	for n, v := range env {
		newEnv[n] = v
	}
	newEnv[name] = value
	return newEnv
}

func DoInfraDeploy(t *testing.T, region string) {

	// Unique run identifier
	uniqueId := random.UniqueId()
	logger.Logf(t, "uniqueId: %v", uniqueId)

	// Temporary terraform data dir
	dataDir, err := ioutil.TempDir("", fmt.Sprintf("terratest-%s-*", uniqueId))
	require.NoError(t, err, "Error creating data directory: %v", err)

	defer os.RemoveAll(dataDir)
	logger.Logf(t, "Data directory: %v", dataDir)

	// Absolute terraform data path
	dataPath, err := filepath.Abs("./.terraform")
	if err != nil {
		t.Fatalf("Error converting file path: %v", err)
	}
	logger.Logf(t, "Data path: %v", dataPath)

	envVars := map[string]string{
		"AWS_REGION": region,
		"TF_LOG":     "TRACE",
		// "TF_LOG_PATH":       fmt.Sprintf("%s/terratest-%s.log", dataPath, uniqueId),
		"TF_INPUT":      "0",
		"TF_VAR_region": region,
		// "TF_DATA_DIR":       fmt.Sprintf("%s/terratest-%s", dataPath, uniqueId),
		"TF_IN_AUTOMATION":  "YES",
		"TF_CLI_ARGS_plan":  "-parallelism=25",
		"TF_CLI_ARGS_apply": "-parallelism=25",
	}

	// Base VPC+SUBNET
	vpc001EnvVars := EnvWithNewVar(envVars, "TF_DATA_DIR", fmt.Sprintf("%s/terratest-%s-vpc-001", dataPath, uniqueId))
	vpc001EnvVars = EnvWithNewVar(vpc001EnvVars, "TF_LOG_PATH", fmt.Sprintf("%s/terratest-%s-vpc-001.log", dataPath, uniqueId))
	vpc_001 := &terraform.Options{
		TerraformDir: "../.infratest",
		Vars: map[string]interface{}{
			"vpc_name": fmt.Sprintf("terratest-%s", uniqueId),
		},
		EnvVars: vpc001EnvVars,
	}
	terraform.WorkspaceSelectOrNew(t, vpc_001, fmt.Sprintf("terratest-%s", uniqueId))
	defer terraform.Destroy(t, vpc_001)

	terraform.Init(t, vpc_001)
	terraform.Plan(t, vpc_001)
	terraform.ApplyAndIdempotent(t, vpc_001)

	logger.Logf(t, "vpc_id: %s", terraform.OutputRequired(t, vpc_001, "vpc_id"))
	logger.Logf(t, "az: %s", terraform.OutputRequired(t, vpc_001, "az"))

	// Deploy module directly
	subnet001EnvVars := EnvWithNewVar(envVars, "TF_DATA_DIR", fmt.Sprintf("%s/terratest-%s-subnet-001", dataPath, uniqueId))
	subnet001EnvVars = EnvWithNewVar(subnet001EnvVars, "TF_LOG_PATH", fmt.Sprintf("%s/terratest-%s-subnet-001.log", dataPath, uniqueId))
	subnet001 := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"cidr":   "10.0.2.0/24",
			"vpc_id": terraform.OutputRequired(t, vpc_001, "vpc_id"),
			"az":     terraform.OutputRequired(t, vpc_001, "az"),
			"name":   fmt.Sprintf("terratest-%s", uniqueId),
		},
		EnvVars: subnet001EnvVars,
	}
	terraform.WorkspaceSelectOrNew(t, subnet001, fmt.Sprintf("terratest-%s", uniqueId))
	defer terraform.Destroy(t, subnet001)

	terraform.Init(t, subnet001)
	terraform.Plan(t, subnet001)
	terraform.ApplyAndIdempotent(t, subnet001)

	logger.Logf(t, "subnet_id: %v", terraform.OutputRequired(t, subnet001, "subnet_id"))
	logger.Logf(t, "rt_id: %v", terraform.OutputRequired(t, subnet001, "rt_id"))

}

func TestRegionFromEnv(t *testing.T) {
	t.Parallel()
	region, ok := os.LookupEnv("TERRATEST_REGION")
	require.True(t, ok, "Variable is required: %v", "TERRATEST_REGION")
	require.NotEmpty(t, region, "Variable should not be empty: %v", "TERRATEST_REGION")
	logger.Logf(t, "Testing region: %v", region)
	DoInfraDeploy(t, region)
}
