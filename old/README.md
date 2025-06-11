# Calvin 🦦 – Your Smart Customer Manager  

Calvin is your new customer management assistant!  
Easily manage your customers, create cost estimates, track sales, and much more – all in one place.  

## 📌 Project Releases  

### 🔹 Version 0.1  
✅ Local functionality for managing customer information  
✅ Save and retrieve customer details  
❌ No database integration yet (coming soon!)  

## Running with PostgreSQL

This project now uses PostgreSQL as its database to store and manage application data. The setup is managed using Docker Compose, which simplifies the process of starting both the API and the database.

### Getting Started

1.  **Start the services:**
    Open your terminal, navigate to the project's root directory, and run the following command:
    ```bash
    docker-compose up -d
    ```
    This command will build the API image (if not already built) and start both the API service (`api`) and the PostgreSQL database service (`postgres_db`) in detached mode. The API service is configured to wait until the database service is healthy before starting.

2.  **Database Connection:**
    The API service connects to the PostgreSQL database using the `DB_DSN` (Database Source Name) environment variable. This is pre-configured in the `docker-compose.yml` file for the `api` service to connect to the `postgres_db` service.

    The default connection details used are:
    *   **User:** `calvinuser`
    *   **Password:** `calvinpass` (as set in your `.env` file, though `docker-compose.yml` directly uses `calvinpass` for `POSTGRES_PASSWORD`)
    *   **Database Name:** `calvindb`
    *   **Host:** `postgres_db` (service name within the Docker network)
    *   **Port:** `5432`

    If you need to change these credentials or the database name, you can modify them in the `environment` section of the `postgres_db` service and update the `DB_DSN` for the `api` service within the `docker-compose.yml` file.

3.  **Accessing the Database Directly (Optional):**
    If you need to connect directly to the PostgreSQL database (e.g., for debugging or manual queries), you can use `psql` via Docker. Once the services are running, execute the following command in your terminal:
    ```bash
    docker exec -it postgres_db_calvin psql -U calvinuser -d calvindb
    ```
    You will be prompted for the password, which is `calvinpass`.

With these steps, the API should be running and connected to the PostgreSQL database.
