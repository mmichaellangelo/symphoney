FROM alpine:3.21

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache \
    nodejs \
    npm \
    nginx

# Copy node package files
COPY package*.json ./

# Install node dependencies
RUN npm install

# Copy rest of source files
COPY . .

# Build static site
RUN npm run build

# Set permissions for build directory
RUN chown -R nginx:nginx /app/build
RUN chmod -R 755 /app/build

# Copy nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port for HTTP (HTTPS TODO)
EXPOSE 80

# Start the server
CMD ["nginx", "-g", "daemon off;"]
