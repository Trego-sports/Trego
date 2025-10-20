# ────────────────
# Stage 1: Build
# ────────────────
FROM oven/bun:1 AS builder

# Set working directory
WORKDIR /app

# Copy only necessary files for dependency resolution
COPY bun.lock package.json ./

# Install dependencies using Bun (fast & cache-friendly)
RUN bun install --frozen-lockfile

# Copy the rest of the project files
COPY . .

# Build the production-ready app (Vite build)
RUN bun run build

# ────────────────────────────────
# Stage 2: Serve (Static Nginx)
# ────────────────────────────────
FROM nginx:alpine

# Copy built files from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy optional Nginx config (uncomment if custom routes)
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
