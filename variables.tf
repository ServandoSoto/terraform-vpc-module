variable "vpc_cidr" {
  description = "CIDR de la VPC"
  type        = string
}

variable "vpc_name" {
  description = "Nombre de la VPC"
  type        = string
}

variable "public_subnet_cidr" {
  description = "CIDR del subnet público"
  type        = string
}   

variable "private_subnet_cidr" {
  description = "CIDR del subnet privado"
  type        = string
}

variable "az" {
  description = "Zona de disponibilidad"
  type        = string
}  

variable "security_group_name" {
  description = "Nombre del grupo de seguridad"
  type        = string
}