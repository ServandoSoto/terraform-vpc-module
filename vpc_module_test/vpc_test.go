// Paquete test — nombre estándar para tests en Go
package test

import (
	"testing" // Librería estándar de Go para tests

	// Terratest: librería que permite desplegar infraestructura Terraform real y validarla
	"github.com/gruntwork-io/terratest/modules/terraform"
	// Testify: librería para hacer assertions (verificar que algo cumple una condición)
	"github.com/stretchr/testify/assert"
)

// TestVPCModule es la función principal del test — Go la reconoce como test por empezar con "Test"
func TestVPCModule(t *testing.T) {

	// terraformOptions define la configuración del despliegue:
	// TerraformDir apunta al módulo que vamos a testear
	// Vars son los valores que le pasamos a las variables del módulo
	terraformOptions := &terraform.Options{
		TerraformDir: "../vpc_module",
		Vars: map[string]interface{}{
			"vpc_cidr":            "10.0.0.0/16",
			"vpc_name":            "vpc-test",
			"public_subnet_cidr":  "10.0.1.0/24",
			"private_subnet_cidr": "10.0.2.0/24",
			"az":                  "us-east-1a",
			"security_group_name": "web-sg-test",
		},
	}

	// defer garantiza que terraform destroy se ejecute al final del test
	// aunque el test falle — evita dejar recursos huérfanos en AWS
	defer terraform.Destroy(t, terraformOptions)

	// InitAndApply ejecuta terraform init + terraform apply en el módulo
	// despliega la infraestructura real en AWS
	terraform.InitAndApply(t, terraformOptions)

	// terraform.Output captura el valor del output del módulo
	// assert.NotEmpty verifica que no esté vacío — si lo está, el test falla
	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	assert.NotEmpty(t, vpcID)

	publicSubnetID := terraform.Output(t, terraformOptions, "public_subnet_id")
	assert.NotEmpty(t, publicSubnetID)

	privateSubnetID := terraform.Output(t, terraformOptions, "private_subnet_id")
	assert.NotEmpty(t, privateSubnetID)

	sgID := terraform.Output(t, terraformOptions, "security_group_id")
	assert.NotEmpty(t, sgID)
}
