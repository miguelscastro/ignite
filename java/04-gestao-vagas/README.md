# Gestão de Vagas — API

Job management API built with Spring Boot. This project includes Docker Compose for local services and Terraform configuration for cloud/infrastructure examples.

---

## 🔧 Prerequisites
- Java 11+ or 17+
- Docker & Docker Compose
- Maven (or use the included wrapper `./mvnw`)

---

## 🚀 Quick start (local)

Start the stack with Docker Compose:

```bash
cd java/04-gestao-vagas
./mvnw clean package
# start DB and app via compose
docker-compose up --build
```

Or run the app directly with the maven wrapper:

```bash
./mvnw spring-boot:run
```

---

## ⚙️ Notes
- Prometheus configuration is available in `config/prometheus.yml`.
- Terraform files are in `terraform/` — review `terraform.tfvars` before applying.
- Ports and environment variables are defined in `docker-compose.yml` and `src/main/resources`.

---

Happy developing! ✅