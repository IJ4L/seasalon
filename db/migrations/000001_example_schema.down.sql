-- Drop triggers
DROP TRIGGER IF EXISTS update_users_updated_at_trigger ON users;
DROP TRIGGER IF EXISTS update_branch_updated_at_trigger ON branches;
DROP TRIGGER IF EXISTS update_services_updated_at_trigger ON services;
DROP TRIGGER IF EXISTS update_reservations_updated_at_trigger ON reservations;
DROP TRIGGER IF EXISTS update_rating_updated_at_trigger ON ratings;

-- Drop trigger functions
DROP FUNCTION IF EXISTS update_users_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_branch_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_services_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_reservations_updated_at CASCADE;
DROP FUNCTION IF EXISTS update_rating_updated_at CASCADE;

-- Drop foreign key constraints
ALTER TABLE reservations DROP CONSTRAINT IF EXISTS reservations_iduser_fkey;
ALTER TABLE reservations DROP CONSTRAINT IF EXISTS reservations_idservice_fkey;
ALTER TABLE ratings DROP CONSTRAINT IF EXISTS ratings_iduser_fkey;
ALTER TABLE ratings DROP CONSTRAINT IF EXISTS ratings_idservice_fkey;
ALTER TABLE services DROP CONSTRAINT IF EXISTS services_idbranch_fkey;

-- Drop tables
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS branches;
DROP TABLE IF EXISTS users;
