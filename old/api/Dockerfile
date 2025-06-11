# Use an official Node.js runtime as a parent image
FROM node:14

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in package.json
RUN npm install

# Make port 8000 available to the world outside this container
EXPOSE 8000

# Run the command to start your application
CMD ["sh", "-c", "./docs/serve.js"]