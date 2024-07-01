-- Drop triggers
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP TRIGGER IF EXISTS update_branch_updated_at ON Branch;
DROP TRIGGER IF EXISTS update_services_updated_at ON services;
DROP TRIGGER IF EXISTS update_reservations_updated_at ON Reservations;
DROP TRIGGER IF EXISTS update_rating_updated_at ON Rating;

-- Drop trigger functions
DROP FUNCTION IF EXISTS update_updated_at_column CASCADE;
DROP FUNCTION IF EXISTS update_branch_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_services_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_reservations_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_rating_updated_at CASCADE;

-- Drop foreign key constraints
ALTER TABLE Reservations DROP CONSTRAINT IF EXISTS reservations_iduser_fkey;
ALTER TABLE Reservations DROP CONSTRAINT IF EXISTS reservations_idservice_fkey;
ALTER TABLE Rating DROP CONSTRAINT IF EXISTS rating_iduser_fkey;
ALTER TABLE Rating DROP CONSTRAINT IF EXISTS rating_idservice_fkey;
ALTER TABLE services DROP CONSTRAINT IF EXISTS services_idbranch_fkey;

-- Drop tables
DROP TABLE IF EXISTS Rating;
DROP TABLE IF EXISTS Reservations;
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS Branch;
DROP TABLE IF EXISTS users;
