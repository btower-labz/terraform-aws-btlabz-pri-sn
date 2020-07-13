package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"os"
	"path/filepath"
	"testing"
)

func DoInfraDeploy(t *testing.T, region string) {
	uniqueId := random.UniqueId()
	logger.Logf(t, "uniqueId: %v", uniqueId)

	dataPath, err := filepath.Abs("./.terraform")
	if err != nil {
		t.Fatalf("Error converting file path: %v", err)
	}

	os.MkdirAll(dataPath, 0755)

	envVars := map[string]string{
		"AWS_REGION":        region,
		"TF_LOG":            "TRACE",
		"TF_LOG_PATH":       fmt.Sprintf("%s/terratest-%s.log", dataPath, uniqueId),
		"TF_INPUT":          "0",
		"TF_VAR_region":     region,
		"TF_DATA_DIR":       fmt.Sprintf("%s/terratest-%s", dataPath, uniqueId),
		"TF_IN_AUTOMATION":  "YES",
		"TF_CLI_ARGS_plan":  "-parallelism=25",
		"TF_CLI_ARGS_apply": "-parallelism=25",
	}

	vpc_001 := &terraform.Options{
		TerraformDir: "../.infratest",
		Vars: map[string]interface{}{
			"vpc_name": fmt.Sprintf("terratest-%s", uniqueId),
		},
		EnvVars: envVars,
	}
	terraform.WorkspaceSelectOrNew(t, vpc_001, fmt.Sprintf("terratest-%s", uniqueId))
	defer terraform.Destroy(t, vpc_001)

	terraform.Init(t, vpc_001)
	terraform.Plan(t, vpc_001)
	terraform.ApplyAndIdempotent(t, vpc_001)

	logger.Logf(t, "vpc_id: %v", terraform.OutputRequired(t, vpc_001, "vpc_id"))
	logger.Logf(t, "az: %v", terraform.OutputRequired(t, vpc_001, "az"))

        /*

	subnet001 := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"cidr":   "10.0.2.0/24",
			"vpc_id": terraform.OutputRequired(t, vpc_001, "vpc_id"),
			"az":     terraform.OutputRequired(t, vpc_001, "az"),
			"name":   fmt.Sprintf("terratest-%s", uniqueId),
		},
		EnvVars: envVars,
	}
	terraform.WorkspaceSelectOrNew(t, subnet001, fmt.Sprintf("terratest-%s", uniqueId))
	defer terraform.Destroy(t, subnet001)

	terraform.Init(t, subnet001)
	terraform.Plan(t, subnet001)
	terraform.ApplyAndIdempotent(t, subnet001)

	logger.Logf(t, "subnet_id: %v", terraform.OutputRequired(t, subnet001, "subnet_id"))
	logger.Logf(t, "rt_id: %v", terraform.OutputRequired(t, subnet001, "rt_id"))

        */

}

func TestRegion_US_WEST_1(t *testing.T) {
	t.Parallel()
	DoInfraDeploy(t, "us-west-1")
}

func TestRegion_EU_WEST_1(t *testing.T) {
	t.Parallel()
	DoInfraDeploy(t, "eu-west-1")
}
