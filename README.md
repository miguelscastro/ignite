# Ignite

**Full Stack technical growth repository** — a curated collection of projects and challenges from the **Ignite** training developed by Faculdade de Tecnologia Rocketseat (Rocketseat Technology College). Find exercises, complete apps, and architectures using **Java _(Spring Boot)_**, **Node.js**, **React _(Vite)_**, **Next.js** and **Go _(Golang)_**.

## 🔗 Index

### Java

- [java/01-fundamentos-java](./java/01-fundamentos-java) — Core Java exercises and examples.
- [java/01-livraria](https://github.com/miguelscastro/livraria) — Bookstore examples (external submodule/link).
- [java/02-fundamentos-java-mavenjdbc](./java/02-fundamentos-java-mavenjdbc) — JDBC + Maven.
- [java/03-fundamentos-springboot](./java/03-fundamentos-springboot) — Spring Boot fundamentals.
- [java/04-gestao-vagas](https://github.com/miguelscastro/gestao-vagas) — Job management API (includes Docker & Prometheus).
- [java/05-gestao-vagas-web](./java/05-gestao-vagas-web) — Web front-end for job management.
- [java/06-petland](./java/06-petland) — Petland project with Docker orchestration.

### Go

- [go/01-fundamentos-go](./go/01-fundamentos-go) — Go basics (packages, functions, variables, arrays, loops, conditionals, defer).
- [go/01-guessing-game](./go/01-guessing-game) — CLI guessing game.
- [go/02-currency-exchange](./go/02-currency-exchange) — Currency converter example.
- [go/02-one-billion-row-challenge](./go/02-one-billion-row-challenge) — Data processing challenge.
- [go/03-bitly](./go/03-bitly) — URL shortener API.
- [go/03-get-movies](./go/03-get-movies) — External API movie search service.
- [go/03-restful-api](./go/03-restful-api) — RESTful API examples (server, JSON, logging).
- [go/04-bitly-redis](https://github.com/miguelscastro/bitly-redis) — Redis-backed URL shortener.
- [go/04-databases](./go/04-databases) — SQL/SQLite and integration examples.
- [go/05-gobid](https://github.com/miguelscastro/go-bid) — GoBid service with Docker Compose.

### Node.js

- [nodejs/01-fundamentos-nodejs](./nodejs/01-fundamentos-nodejs) — Node.js fundamentals and exercises.

### React (reactjs)

- [reactjs/01-fundamentos-reactjs](./reactjs/01-fundamentos-reactjs) — React + Vite (basic).
- [reactjs/01-fundamentos-reactjs-ts](./reactjs/01-fundamentos-reactjs-ts) — React + TypeScript.
- [reactjs/01-todo-list](https://github.com/miguelscastro/to-do-list) — 🎯 **Challenge Project** — To‑do list (submodule/external).
- [reactjs/02-coffee-delivery](https://github.com/miguelscastro/coffee-delivery) — 🎯 **Challenge Project** — Coffee Delivery UI (submodule/external).
- [reactjs/02-ignite-timer](./reactjs/02-ignite-timer) — Timer.
- [reactjs/03-dt-money](./reactjs/03-dt-money) — Finance tracker.
- [reactjs/03-github-blog](https://github.com/miguelscastro/github-blog) — 🎯 **Challenge Project** — GitHub API-based blog (submodule/external).
- [reactjs/04-pizzashop/pizzashop-api](https://github.com/miguelscastro/pizzashop/tree/main/pizzashop-api) — Pizzashop API.
- [reactjs/04-pizzashop/pizzashop-web](https://github.com/miguelscastro/pizzashop/tree/main/pizzashop-web) — Pizzashop front-end.

### Next.js (nextjs)

- [nextjs/01-ignite-shop](./nextjs/01-ignite-shop) — Demo e-commerce built with Next.js.

> Note: Some projects are kept as submodules/external links depending on complexity (some of them marked as **Challenge Project**) — check local README or remote link for details.

---

## 🌐 Active Deploys

| Project                      | Deploy                                             |
| ---------------------------- | -------------------------------------------------- |
| `reactjs/01-todo-list`       | _https://todo-miguelscastro.vercel.app/_           |
| `reactjs/02-coffee-delivery` | _https://coffeedelivery-miguelscastro.vercel.app/_ |
| `reactjs/03-github-blog`     | _https://githubblog-miguelscastro.vercel.app/_     |

---

## 🧰 Prerequisites (short)

- Java 11+ or 17+ (check `pom.xml` per project)
- Maven (`./mvnw` available in most Java projects)
- Node.js 16+ (recommended 18+) and npm / yarn / pnpm
- Docker & Docker Compose (for orchestrated projects)

---

## 🚀 Quick Start (by stack)

### Java / Spring Boot

- Run with the Maven wrapper (preferred):

```bash
cd java/03-fundamentos-springboot
./mvnw spring-boot:run
```

- Build and run JAR:

```bash
./mvnw clean package
java -jar target/*.jar
```

- With Docker Compose (e.g., `04-gestao-vagas`, `06-petland`):

```bash
cd java/04-gestao-vagas
docker-compose up --build
```

### Node.js

- Example:

```bash
cd nodejs/01-fundamentos-nodejs
npm install
npm run dev
```

- If the project uses `pnpm`:

```bash
pnpm install
pnpm dev
```

### React (Vite)

- Example with `reactjs/01-fundamentos-reactjs`:

```bash
cd reactjs/01-fundamentos-reactjs
npm install
npm run dev
```

- Production build:

```bash
npm run build
# serve the dist folder with a static server (e.g., serve)
```

> Projects marked **🎯 Challenge Project** may have specific instructions in their local README; check each folder.

### Next.js

- Example with `nextjs/01-ignite-shop`:

```bash
cd nextjs/01-ignite-shop
pnpm install # or npm install
pnpm dev # or npm run dev
```

- Production build:

```bash
pnpm build
pnpm start
```

### Go

- Example with `go/02-currency-exchange`:

```bash
cd go/02-currency-exchange
go run main.go
```

---

## 🛠️ Notes & Best practices

- Always check the subproject README for environment variables and specific requirements.
- Use wrappers (`./mvnw`, `pnpm`, etc.) for reproducible builds; check `go.mod` for Go dependency versions.
- Keep numeric prefixes (`01-`, `02-`) to preserve chronological order.

---

## 📄 License

See each subproject README or the root LICENSE file for details.
