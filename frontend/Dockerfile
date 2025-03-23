# BUILD STAGE
FROM node:23-alpine as build

# Set working directory
WORKDIR /app

# Copy node package files
COPY package*.json ./

# Install node dependencies
RUN npm install

# Copy rest of source files
COPY . .

# Get API address from build args
ARG API_ADDRESS=symphoney.xyz

# Set environment var
ENV API_ADDRESS=${API_ADDRESS}

# Build static site
RUN npm run build

# Use nginx alpine image as final stage
FROM nginx:1.27.4-alpine

# Copy built webapp to nginx image
COPY --from=build /app/build /app

# Copy nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port for HTTP (HTTPS TODO)
EXPOSE 80

# Start the server
CMD ["nginx", "-g", "daemon off;"]
