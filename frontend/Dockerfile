# Use an official Node.js runtime as base image
FROM node:18.19.0

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire frontend codebase
COPY . .

# Expose the React/Next.js port
EXPOSE 5173

# Start the frontend app
CMD ["npm", "run", "dev"]
