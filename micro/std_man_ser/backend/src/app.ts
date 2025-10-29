import express from "express";
import cors from "cors";
import bodyParser from "body-parser";
import swaggerUi from "swagger-ui-express";
import swaggerJsdoc from "swagger-jsdoc";
import routes from "./routes";

const app = express();

app.use(cors());
app.use(bodyParser.json());

// Swagger setup (placeholder - update docs/openapi.yaml as needed)
const swaggerSpec = swaggerJsdoc({
  definition: {
    openapi: "3.0.0",
    info: {
      title: "Student Management Service API",
      version: "0.1.0",
      description:
        "APIs for managing students, colleges, departments, and programs",
    },
  },
  apis: ["./src/routes/*.ts", "./src/controllers/*.ts"],
});

app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerSpec));

app.use("/api", routes);

// basic 404
app.use((req, res) => {
  res
    .status(404)
    .json({ ok: false, error: { code: "not_found", message: "Not found" } });
});

// error handler
app.use((err: any, req: any, res: any, next: any) => {
  // eslint-disable-next-line no-console
  console.error(err);
  const status = err.status || 500;
  res
    .status(status)
    .json({
      ok: false,
      error: {
        code: err.code || "internal_error",
        message: err.message || "Internal server error",
      },
    });
});

export default app;
