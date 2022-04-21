package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(test *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(test, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../examples",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(test, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(test, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(test, terraformOptions, "hello_world")
	assert.Equal(test, "Hello, World!", output)
}
