# Ignite — Monorepo

A curated collection of projects built during the **Ignite** formation modules. This repository groups Java backend projects (Spring Boot) and React frontend applications. Use this README as an index and quick start guide for each project.

---

## 🔗 Index

### Java
- [01-fundamentos-java](java/01-fundamentos-java/) — Core/essential Java exercises and examples.
- [01-livraria](java/01-livraria/) — Bookstore examples and exercises.
- [02-fundamentos-java-mavenjdbc](java/02-fundamentos-java-mavenjdbc/) — JDBC + Maven exercises.
- [03-fundamentos-springboot](java/03-fundamentos-springboot/) — Spring Boot fundamentals.
- [04-gestao-vagas](java/04-gestao-vagas/) — Job management API; includes Docker and Prometheus config.
- [05-gestao-vagas-web](java/05-gestao-vagas-web/) — Web front-end for job management.
- [06-petland](java/06-petland/) — Petland project (contains Docker setup and services).

### React
- [01-fundamentos-reactjs](react/01-fundamentos-reactjs/) — React + Vite basics.
- [01-fundamentos-reactjs-ts](react/01-fundamentos-reactjs-ts/) — TypeScript + React fundamentals.
- [01-todo-list](react/01-todo-list/) — Simple todo list app.
- [02-coffee-delivery](react/02-coffee-delivery/) — Coffee shop UI exercises.
- [02-ignite-timer](react/02-ignite-timer/) — Timer app.
- [03-dt-money](react/03-dt-money/) — Finance tracker.
- [03-github-blog](react/03-github-blog/) — GitHub API-based blog.
- [04-pizzashop/pizzashop-api](react/04-pizzashop/pizzashop-api/) — Pizza shop API.
- [04-pizzashop/pizzashop-web](react/04-pizzashop/pizzashop-web/) — Pizza shop front-end.

> Note: Each link points to the project folder; many subprojects include their own README with more details.

---

## 🧰 Prerequisites (high level)
- Java 11+ or 17+ (projects may target different versions; check the subproject pom.xml)
- Maven (or use the provided wrapper `./mvnw` on Linux)
- Node.js 16+ and npm (or yarn/pnpm as preferred)
- Docker & Docker Compose (for projects that include `docker-compose.yml`)

---

## 🚀 Quick start (common commands)

### Java / Spring Boot
General options when a `pom.xml` is present:

- Run with the Maven wrapper (preferred):

```bash
cd java/03-fundamentos-springboot
./mvnw spring-boot:run
```

- Build and run the packaged JAR:

```bash
./mvnw clean package
java -jar target/*.jar
```

- If a project provides Docker Compose (e.g., `04-gestao-vagas`, `06-petland`):

```bash
cd java/04-gestao-vagas
docker-compose up --build
```

Ports and DBs:
- Common API port: `8080` (verify each project for exact configuration)
- PostgreSQL (if used) usually maps to `5432` in compose setups — check each project for details.

### React / Vite
For front-end projects with a `package.json`:

```bash
cd react/01-fundamentos-reactjs
npm install
npm run dev
```

- Default dev port for Vite: `5173` (or as printed in the terminal).
- Production build: `npm run build` and serve the `dist/` folder with a static server.

---

## 🛠️ Notes & tips
- Prefer the project-local README when available — it contains project-specific requirements and environment variables.
- Use the included Maven wrapper (`./mvnw`) instead of a globally installed Maven to ensure consistent builds.
- If a project requires a DB or external service, look for `docker-compose.yml` or `README.md` inside that subfolder.

---

## 📬 Contributing
- Add new projects under `java/` or `react/` with a short README explaining how to run them.
- Keep folder names prefixed with an incremental number for chronological order (e.g., `01-`, `02-`).

---

## 📄 License
See the repository root or individual project READMEs for license details.
