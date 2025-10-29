# Student Management Service (backend)

This folder contains the TypeScript + Express backend using Prisma (SQLite) for the Student Management Service.

Quick start

1. Copy .env file and set DATABASE_URL (sqlite):

```bash
cp .env.example .env
# edit .env to configure DATABASE_URL, e.g. DATABASE_URL="file:./dev.db"
```

2. Install dependencies and generate Prisma client:

```bash
cd micro/std_man_ser/backend
npm install
npx prisma generate
npx prisma migrate dev --name init
```

3. Run in development:

```bash
npm run dev
```

Notes

- APIs are under `/api` by default. Swagger UI is available at `/api-docs`.
- Tests: `npm test` (Jest)
