# syntax=docker/dockerfile:1.7

############################
# 1️⃣  Build stage — Bun
############################
FROM oven/bun:1 AS builder
WORKDIR /app

# Copy only dependency manifests first (for caching)
COPY package.json bun.lock* ./

# Install dependencies
RUN bun install --frozen-lockfile

# Copy the rest of your project
COPY . .

# Build the Vite app
RUN bun run build



############################
# 2️⃣  Run stage — Nginx
############################
FROM nginx:1.27-alpine AS run

# Replace default Nginx site config with our SPA config
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy built assets from Bun build stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose web port
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
