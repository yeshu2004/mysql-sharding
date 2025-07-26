# MySQL Sharding Implementation in Go

This project demonstrates a basic sharding implementation in Go using MySQL, where data is distributed across multiple database instances (`first_recording` and `second_recording`) based on a sharding key.

## Overview

- **Sharding Logic**: The `albumsByIdSharded` function implements sharding by using the modulo operation (`id % 10`) on the album ID. IDs with a modulo less than 5 are stored in `db1`, while those with a modulo between 5 and 9 are stored in `db2`.
- **Database Connections**: The program establishes connections to two MySQL databases using the `github.com/go-sql-driver/mysql` driver and manages them via the `database/sql` package.
- **Environment Configuration**: Database credentials and connection details are loaded from a `.env` file using the `github.com/joho/godotenv` package.

## Setup

1. **Install Dependencies**:
   ```bash
   go get github.com/go-sql-driver/mysql
   go get github.com/joho/godotenv
   ```

2. **Configure Environment**:
   - Create a `.env` file in the project root with the following variables:
     ```
     DBUSER=your_username
     DBPASS=your_password
     ```
   - Ensure MySQL is running locally on `127.0.0.1:3306` with the databases `first_recording` and `second_recording` created.

3. **Run the Application**:
   ```bash
   go run main.go
   ```

## Usage

- **Query by ID**: The `albumsByIdSharded` function retrieves an album by its ID, routing the query to the appropriate database based on the sharding logic.
- **Query by Artist**: The `albumsByArtist` function queries all albums by a given artist from `db1` (non-sharded for simplicity).

## Sharding Details

- **Sharding Key**: The album `ID` is used as the sharding key.
- **Distribution**: 
  - `ID % 10 < 5` → `db1`
  - `5 <= ID % 10 < 10` → `db2`
- **Limitations**: This is a basic implementation with a fixed modulo range. For production, consider dynamic sharding keys and more robust distribution strategies.

## Future Improvements

- Add support for additional databases.
- Implement dynamic sharding key configuration.
- Enhance error handling and connection pooling.