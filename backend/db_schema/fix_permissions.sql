-- Fix Database Permissions Script
-- This script grants necessary permissions to the database user for sequences and tables
-- Run this script as a PostgreSQL superuser (usually 'postgres')

-- Note: Replace 'your_db_user' with your actual database username
-- You can find your username in your .env file or environment variables

-- Grant usage on all sequences in the current database
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO your_db_user;

-- Grant permissions on existing tables
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_db_user;

-- Grant permissions on future sequences (for new tables)
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT ON SEQUENCES TO your_db_user;

-- Grant permissions on future tables
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO your_db_user;

-- Specific grants for the components sequence (the one causing the error)
GRANT USAGE, SELECT ON SEQUENCE components_id_seq TO your_db_user;
GRANT USAGE, SELECT ON SEQUENCE retailers_id_seq TO your_db_user;
GRANT USAGE, SELECT ON SEQUENCE prices_id_seq TO your_db_user;
GRANT USAGE, SELECT ON SEQUENCE user_builds_id_seq TO your_db_user;
GRANT USAGE, SELECT ON SEQUENCE build_components_id_seq TO your_db_user;

-- Grant permissions on the tables themselves
GRANT ALL PRIVILEGES ON TABLE components TO your_db_user;
GRANT ALL PRIVILEGES ON TABLE retailers TO your_db_user;
GRANT ALL PRIVILEGES ON TABLE prices TO your_db_user;
GRANT ALL PRIVILEGES ON TABLE user_builds TO your_db_user;
GRANT ALL PRIVILEGES ON TABLE build_components TO your_db_user;

-- Grant schema usage permission
GRANT USAGE ON SCHEMA public TO your_db_user;

-- Optional: Make the user owner of the database (uncomment if needed)
-- ALTER DATABASE your_database_name OWNER TO your_db_user;
