# Go MySQL Sharding Example

This project demonstrates a simple sharding approach in Go using two MySQL databases. Album records are distributed between two databases (`first_recording` and `secound_recording`) based on a hash of their UUID.

## Features

- **Sharding Logic:** Uses FNV-1a hash on the album UUID. If `hash % 2 == 0`, the record goes to `db1` (`first_recording`); otherwise, it goes to `db2` (`secound_recording`).
- **CRUD Operations:** 
  - Add a new album (`addAlbum`)
  - Get album by ID (`getAlbum`)
  - Query albums by artist across both shards (`albumsByArtist`)
- **Environment Configuration:** Loads DB credentials from a `.env` file.

## Setup

1. **Install Dependencies**
   ```bash
   go get github.com/go-sql-driver/mysql
   go get github.com/joho/godotenv
   go get github.com/google/uuid
   ```

2. **Configure MySQL**
   - Create two databases: `first_recording` and `secound_recording`.
   - Create the `album` table in both databases:
     ```sql
     CREATE TABLE album (
         id VARCHAR(36) NOT NULL,
         title VARCHAR(128) NOT NULL,
         artist VARCHAR(128) NOT NULL,
         price DECIMAL(5,2) NOT NULL,
         PRIMARY KEY (id)
     );
     ```

3. **Create a `.env` File**
   ```
   DBUSER=your_mysql_user
   DBPASS=your_mysql_password
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```

## How Sharding Works

- The function `getShardById(id string)` hashes the UUID and uses `%2` to select the database.
- This ensures that the same UUID always maps to the same shard.

## Example Usage

- **Add an Album:**  
  Adds a new album to the correct shard based on its generated UUID.
- **Get Album by ID:**  
  Retrieves an album from the correct shard using its UUID.
- **Get Albums by Artist:**  
  Searches both shards for albums by a given artist.

## Notes

- This is a basic sharding demo. For production, consider more shards and advanced hash/distribution strategies.
- The `.env` file should be added to `.gitignore` to avoid exposing credentials.

## License