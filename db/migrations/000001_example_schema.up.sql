-- Tabel users
CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    Fullname VARCHAR(255) NOT NULL,
    Photo VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    PhoneNumber VARCHAR(20) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Role VARCHAR(10) CHECK (Role IN ('admin', 'customer')) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Fungsi trigger untuk mengupdate UpdatedAt pada tabel users
CREATE OR REPLACE FUNCTION update_users_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Trigger untuk memanggil fungsi update_users_updated_at sebelum UPDATE pada tabel users
CREATE TRIGGER update_users_updated_at_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_users_updated_at();

-- Tabel branches
CREATE TABLE IF NOT EXISTS branches (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Location VARCHAR(255) NOT NULL,
    OpeningTime TIME NOT NULL,
    ClosingTime TIME NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Fungsi trigger untuk mengupdate UpdatedAt pada tabel branches
CREATE OR REPLACE FUNCTION update_branch_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Trigger untuk memanggil fungsi update_branch_updated_at sebelum UPDATE pada tabel branches
CREATE TRIGGER update_branch_updated_at_trigger
BEFORE UPDATE ON branches
FOR EACH ROW
EXECUTE FUNCTION update_branch_updated_at();

-- Tabel services
CREATE TABLE IF NOT EXISTS services (
    ID SERIAL PRIMARY KEY,
    IDBranch INT NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Type VARCHAR(50) NOT NULL,
    Detail TEXT NOT NULL,
    Pricing DECIMAL(10, 2) NOT NULL,
    Duration INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (IDBranch) REFERENCES branches(ID)
);

-- Fungsi trigger untuk mengupdate UpdatedAt pada tabel services
CREATE OR REPLACE FUNCTION update_services_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Trigger untuk memanggil fungsi update_services_updated_at sebelum UPDATE pada tabel services
CREATE TRIGGER update_services_updated_at_trigger
BEFORE UPDATE ON services
FOR EACH ROW
EXECUTE FUNCTION update_services_updated_at();

-- Tabel reservations
CREATE TABLE IF NOT EXISTS reservations (
    ID SERIAL PRIMARY KEY,
    IDUser INT NOT NULL,
    IDService INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (IDUser) REFERENCES users(ID),
    FOREIGN KEY (IDService) REFERENCES services(ID)
);

-- Fungsi trigger untuk mengupdate UpdatedAt pada tabel reservations
CREATE OR REPLACE FUNCTION update_reservations_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Trigger untuk memanggil fungsi update_reservations_updated_at sebelum UPDATE pada tabel reservations
CREATE TRIGGER update_reservations_updated_at_trigger
BEFORE UPDATE ON reservations
FOR EACH ROW
EXECUTE FUNCTION update_reservations_updated_at();

-- Tabel ratings
CREATE TABLE IF NOT EXISTS ratings (  
    ID SERIAL PRIMARY KEY,
    IDUser INT NOT NULL,
    IDService INT NOT NULL,
    Rating INT CHECK (Rating >= 1 AND Rating <= 5) NOT NULL,
    Comment TEXT,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (IDUser) REFERENCES users(ID),
    FOREIGN KEY (IDService) REFERENCES services(ID)
);

-- Fungsi trigger untuk mengupdate UpdatedAt pada tabel ratings
CREATE OR REPLACE FUNCTION update_rating_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Trigger untuk memanggil fungsi update_rating_updated_at sebelum UPDATE pada tabel ratings
CREATE TRIGGER update_rating_updated_at_trigger
BEFORE UPDATE ON ratings
FOR EACH ROW
EXECUTE FUNCTION update_rating_updated_at();
