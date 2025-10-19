# Use Node 20+ (required for React Router v7 and Vite v6)
FROM node:20-alpine

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

# Expose Vite's default dev port
EXPOSE 5173

# Run Vite dev server with proper host binding
CMD ["yarn", "dev", "--host"]
