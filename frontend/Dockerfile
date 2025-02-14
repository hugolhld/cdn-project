# # Use an official Node.js runtime as base image
# FROM node:20.17.0

# # Set working directory
# WORKDIR /app

# # Copy package.json and package-lock.json
# COPY package*.json ./

# # Install dependencies
# RUN npm install

# # Copy the entire frontend codebase
# COPY . .

# # Expose the React/Next.js port
# EXPOSE 5173

# # Start the frontend app
# CMD ["npm", "run", "dev"]


# Étape 1 : Construire le frontend avec Node.js
FROM node:20.17.0 AS builder
WORKDIR /app

# Copier les fichiers package.json et installer les dépendances
COPY package*.json ./
RUN npm install

# Copier le reste du code et builder
COPY . .
RUN npm run build

# Étape 2 : Servir le frontend avec Nginx
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exposer le port 80 (standard HTTP)
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]
