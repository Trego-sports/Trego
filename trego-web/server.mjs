// trego-web/server.mjs
import express from "express";
import { createProxyMiddleware } from "http-proxy-middleware";
import path from "path";
import { fileURLToPath } from "url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const app = express();

// Change only the HOST port in compose if needed; container keeps 3000
const PORT = process.env.PORT || 3000;

// Internal Docker DNS name for your Go service
const API_TARGET = process.env.API_TARGET || "http://trego-backend:8080";

// 1) Serve built static files
app.use(express.static(path.join(__dirname, "dist")));

// 2) Proxy API and health to the Go backend
app.use(
	["/api", "/health"],
	createProxyMiddleware({
		target: API_TARGET,
		changeOrigin: true,
		xfwd: true,
		logLevel: "warn",
	})
);

app.use((_req, res) => {
	res.sendFile(path.join(__dirname, "dist", "index.html"));
});

// 4) Start server
app.listen(PORT, () => {
	console.log(`Web server: http://localhost:${PORT}  (proxy -> ${API_TARGET})`);
});
