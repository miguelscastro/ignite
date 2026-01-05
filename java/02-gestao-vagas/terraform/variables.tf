variable "db_name" {
  description = "database name for the Postgres instance"
  type        = string
  sensitive   = true # Isso impede que a senha apareça nos logs do terminal
}
variable "db_username" {
  description = "superuser for the Postgres instance"
  type        = string
  sensitive   = true # Isso impede que a senha apareça nos logs do terminal
}
variable "db_password" {
  description = "super secret password for the Postgres superuser"
  type        = string
  sensitive   = true # Isso impede que a senha apareça nos logs do terminal
}