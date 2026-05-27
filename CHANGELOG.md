# CHANGELOG - terraform-vpc-module

Todos los cambios notables de este módulo serán documentados en este archivo.
Formato basado en [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [v0.1.0] - 2026-05-26
### Added
- Commit inicial del módulo vpc
- Recurso `aws_vpc` parametrizado con variables
- Recursos `aws_subnet` pública y privada
- Recurso `aws_internet_gateway`
- Recurso `aws_route_table` y `aws_route_table_association`
- Recurso `aws_security_group` con reglas SSH y HTTP
- Archivo `variables.tf` con variables: vpc_cidr, vpc_name, public_subnet_cidr, private_subnet_cidr, az, security_group_name
- Archivo `outputs.tf` con outputs: vpc_id, public_subnet_id, private_subnet_id, security_group_id

---

## [v0.2.0] - 2026-05-26
### Added
- Documentación automática generada con terraform-docs
- Archivo `README.md` con descripción de inputs, outputs y recursos del módulo

---

## [v0.3.0] - 2026-05-26
### Added
- Carpeta `vpc_module_test/` con test de infraestructura usando Terratest
- Archivo `vpc_test.go` que valida outputs vpc_id, public_subnet_id, private_subnet_id, security_group_id
- Archivo `.gitignore` para excluir archivos terraform generados localmente

---

## [v0.3.1] - 2026-05-26
### Fixed
- Corrección de ruta `TerraformDir` en `vpc_test.go` de `../vpc_module` a `..`

---

## [v1.0.0] - 2026-05-27
### Changed
- Versión estable y funcional del módulo
- Arquitectura modular completa desacoplada del código monolítico original
