# 1. Configurações Globais
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1" # Certifique-se de que é a mesma região do seu console
}

# 2. Security Group para a EC2 (Acesso Web e SSH)
resource "aws_security_group" "sg_app" {
  name        = "sg_app_gestao_vagas"
  description = "Permitir HTTP e SSH"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Em produção, use seu IP real aqui
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# 3. Security Group para o RDS (Acesso apenas vindo da EC2)
resource "aws_security_group" "sg_db" {
  name        = "sg_db_gestao_vagas"
  description = "Acesso ao Postgres vindo da EC2"

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.sg_app.id] # Libera apenas para quem tem o SG da App
  }
}

# Chave SSH para acessar a instância EC2
# Cria o recurso de Key Pair na AWS usando sua chave pública local
resource "aws_key_pair" "deployer" {
  key_name   = "minha-chave-ssh"
  public_key = file("~/.ssh/my-aws-key.pub") # Ele lê o conteúdo do arquivo que você gerou
}

# 4. Instância EC2
resource "aws_instance" "app_server" {
  ami           = "ami-0ecb62995f68bb549" # Ubuntu 24.04 LTS (Noble Numbat) em us-east-1
  instance_type = "t3.micro"             # Elegível para Free Tier
  
  key_name      = aws_key_pair.deployer.key_name # Vincula a chave à instância

  vpc_security_group_ids = [aws_security_group.sg_app.id]

  tags = {
    Name = "GestaoVagas-App"
  }
}

# 5. Instância RDS (Postgres)
resource "aws_db_instance" "postgres_db" {
  allocated_storage    = 20
  db_name              = var.db_name
  engine               = "postgres"
  engine_version       = "16"
  instance_class       = "db.t4g.micro" # Free Tier
  username             = var.db_username  
  password             = var.db_password  
  parameter_group_name = "default.postgres16"
  
  vpc_security_group_ids = [aws_security_group.sg_db.id]
  publicly_accessible    = false # Segurança: o banco não fica exposto na internet
  skip_final_snapshot    = true  # Essencial para conseguir dar 'destroy' sem erro
}

# 6. Outputs para facilitar sua vida
output "ec2_public_ip" {
  value = aws_instance.app_server.public_ip
}

output "rds_endpoint" {
  value = aws_db_instance.postgres_db.endpoint
}