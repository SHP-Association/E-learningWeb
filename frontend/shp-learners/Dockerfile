# Use Node 20+ (required for React Router v7 and Vite v6)
FROM node:20-alpine

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install all dependencies (including devDependencies)
RUN npm install

# Copy the rest of the source code
COPY . .

# Expose Vite's default dev port
EXPOSE 5173

# Run Vite dev server with proper host binding
CMD ["npm", "run", "dev", "--", "--host"]
