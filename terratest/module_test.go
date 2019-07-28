package test

import (
	"testing"
	"strings"
	"regexp"
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)


func TestVpcCidr(t *testing.T) {
	
	terraformOptions := &terraform.Options{
		TerraformDir: "./",
		VarFiles:     []string{"module_test.tfvars"},
	}
	
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	cidrBlock := terraform.Output(t, terraformOptions, "cidr_block")
	fmt.Print(cidrBlock)
	assert.True(t, validateVPCCidr(cidrBlock))
}


func validateVPCCidr(cidr string) bool {
	cidr = strings.Trim(cidr, " ")
	fmt.Print(cidr)
	re, _ := regexp.Compile(`(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)`)

	if re.MatchString(cidr) {
			return true
	}

	return false
}
