# Database Permission Fix

## Problem
You're encountering the error: `permission denied for sequence components_id_seq`

This happens when your database user doesn't have the necessary permissions to use sequences that are automatically created with `BIGSERIAL` columns.

## Solution

### Option 1: Quick Fix (Recommended)

1. **Find your database username**
   - Check your `.env` file in the `backend/` directory for `DB_USER`
   - Or check your environment variables: `echo $DB_USER`

2. **Run the permission fix script**
   ```bash
   # Replace 'your_actual_username' with your DB_USER value
   # Replace 'your_database_name' with your DB_NAME value

   # First, edit the fix_permissions.sql file to replace placeholders
   sed -i 's/your_db_user/your_actual_username/g' backend/db_schema/fix_permissions.sql

   # Run as postgres superuser
   sudo -u postgres psql -d your_database_name -f backend/db_schema/fix_permissions.sql
   ```

3. **Alternative: Run commands manually**
   ```bash
   # Connect to PostgreSQL as superuser
   sudo -u postgres psql -d your_database_name

   # Then run these commands (replace 'your_username' with your actual DB_USER):
   GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO your_username;
   GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_username;
   GRANT USAGE ON SCHEMA public TO your_username;

   # Exit psql
   \q
   ```

### Option 2: Database Ownership (Alternative)

If you want your user to have full control:

```bash
# Connect as postgres superuser
sudo -u postgres psql

# Make your user the owner of the database
ALTER DATABASE your_database_name OWNER TO your_username;

# Exit
\q
```

### Option 3: Create User with Proper Permissions

If you need to create a new user:

```bash
# Connect as postgres superuser
sudo -u postgres psql

# Create user with necessary permissions
CREATE USER your_username WITH PASSWORD 'your_password';
ALTER USER your_username CREATEDB;
GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_username;

# Connect to your database
\c your_database_name

# Grant permissions on schema and objects
GRANT USAGE ON SCHEMA public TO your_username;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_username;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO your_username;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO your_username;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT ON SEQUENCES TO your_username;

# Exit
\q
```

## Verification

After applying the fix, test the connection:

```bash
# Test with your scraper
cd scraper
python -c "from database.connection import db_manager; print('Success!' if db_manager.test_connection() else 'Failed!')"
```

## Prevention

For future database setups, ensure your user has proper permissions from the start by:

1. Creating the database with your user as owner
2. Running the setup script as your user
3. Or running the permission grants after table creation

## Common Issues

- **"role does not exist"**: Create the user first
- **"database does not exist"**: Create the database first
- **Still getting permission errors**: Make sure you're using the correct username and database name
- **Cannot connect as postgres**: Install PostgreSQL properly or use `sudo systemctl start postgresql`

## Environment Variables Check

Make sure your `.env` file in the `backend/` directory contains:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_SSLMODE=disable
```
