# 🍕 pizza.shop API

Back-end for a food delivery app (in the spirit of iFood/Uber Eats), built with TypeScript, Drizzle and Elysia. The project is designed to be runtime-agnostic and can run on Bun, Node, or other Web Standard API compatible runtimes.

---

## 🔧 Prerequisites
- Node.js (or Bun) and package manager (npm / bun)
- Docker & Docker Compose (for local DB)

---

## 🚀 Quick start (local)

```bash
# install deps (use bun or npm)
# with bun
bun install
# or with npm
npm install

# start local databases
docker compose up -d

# run migrations and seed data
# with bun
bun migrate
bun seed

# start dev server
bun dev
# or with npm
npm run dev
```

The API will start at the port printed in the console. Check `docker-compose.yml` and the project config for exact ports and DB settings.

---

## 🧪 Tests
- End-to-end tests are included for the main features (see test scripts in package.json)

---

## 🔁 Features (summary)
- Register restaurants and manage restaurant profiles
- Restaurant manager sign-in and authentication flows
- Register customers and create orders
- Manage menu items and restaurant evaluations
- Order lifecycle management
- Metrics endpoints for restaurants

---

## 💡 Notes
- Ensure Docker is running before executing migrations/seed that target the DB container.
- Check `package.json` scripts for runtime-specific commands (bun or npm).

---

Happy hacking! ✅