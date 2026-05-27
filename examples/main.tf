module "vpc" {
  source              = "github.com/ServandoSoto/terraform-vpc-module?ref=v1.0.0"
  vpc_cidr            = "10.0.0.0/16"
  vpc_name            = "vpc-ejemplo"
  public_subnet_cidr  = "10.0.1.0/24"
  private_subnet_cidr = "10.0.2.0/24"
  az                  = "us-east-1a"
  security_group_name = "web-sg"
}

output "vpc_id" {
  value = module.vpc.vpc_id
}

output "public_subnet_id" {
  value = module.vpc.public_subnet_id
}
