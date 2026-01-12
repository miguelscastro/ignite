# Petland

Petland is a sample project from the Ignite modules. It includes services and Docker setup to run locally.

---

## 🔧 Prerequisites
- Java 11+ or 17+
- Docker & Docker Compose
- Maven (or use the included wrapper `./mvnw`)

---

## 🚀 Quick start (local)

```bash
cd java/06-petland
# build
./mvnw clean package
# or use docker-compose if provided
docker-compose up --build
```

See `HELP.md` for project-specific notes.

---

## 💡 Notes
- Check `docker-compose.yml` for service definitions and exposed ports.
- Use feature branches and avoid pushing directly to `main` for deployed projects.

Happy developing! ✅