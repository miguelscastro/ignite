
# Gestão de Vagas — Frontend

A simple frontend for the Job Management (Gestão de Vagas) project built with Spring Boot and Thymeleaf. This project serves as the UI that communicates with the backend API (see `java/04-gestao-vagas`).

---

## 🔧 Prerequisites
- Java 11+ installed
- Maven or use the Maven wrapper (`./mvnw`)
- Backend API running (default: http://localhost:8080)

---

## ⚙️ Configuration
Update `src/main/resources/application.properties` for local development:

```properties
server.port=8082
host.api.gestao.vagas=http://localhost:8080
```

Adjust ports/host values as needed and save the file.

---

## 🚀 Quick start (local)

```bash
# from repository root or project folder
cd java/05-gestao-vagas-web
./mvnw clean install
./mvnw spring-boot:run
```

The app will be available at http://localhost:8082 (or the `server.port` you configured).

---

## 💡 Notes
- This frontend expects the backend API to be available at the URL defined in `host.api.gestao.vagas`.
- If the backend runs on a different port, update `application.properties` accordingly.
- Avoid committing changes to `main` if the project is deployed (e.g., Render); create feature branches and validate before merging.

---

Happy developing! ✅




