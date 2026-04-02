# Build stage
FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Install Yarn globally
RUN npm install -g yarn --force

# Copy package.json and yarn.lock
COPY package.json yarn.lock ./

# Install dependencies with Yarn
RUN yarn install

# Copy the rest of the source code
COPY . .

# Build the application
RUN yarn build

# Production stage
FROM nginx:alpine

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy built files from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost/health || exit 1

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
