# Ejemplo de uso - terraform-vpc-module

Este directorio contiene un ejemplo funcional de cómo usar el módulo `vpc_module`.

## ¿Qué despliega este ejemplo?

- Una VPC con CIDR `10.0.0.0/16`
- Una subnet pública y una subnet privada
- Un Internet Gateway
- Una Route Table pública con su asociación
- Un Security Group con reglas para SSH (22) y HTTP (80)

## Uso

```hcl
module "vpc" {
  source              = "github.com/ServandoSoto/terraform-vpc-module?ref=v1.0.0"
  vpc_cidr            = "10.0.0.0/16"
  vpc_name            = "vpc-ejemplo"
  public_subnet_cidr  = "10.0.1.0/24"
  private_subnet_cidr = "10.0.2.0/24"
  az                  = "us-east-1a"
  security_group_name = "web-sg"
}
```

## Outputs disponibles

| Output | Descripción |
|--------|-------------|
| vpc_id | ID de la VPC creada |
| public_subnet_id | ID de la subnet pública |
| private_subnet_id | ID de la subnet privada |
| security_group_id | ID del security group |

## Requisitos

- Terraform >= 1.0.0
- AWS Provider ~> 5.0
- Credenciales AWS configuradas
