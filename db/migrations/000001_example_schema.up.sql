CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
    Fullname VARCHAR(255) NOT NULL,
    Photo VARCHAR(255),
    Email VARCHAR(255) NOT NULL UNIQUE,
    PhoneNumber VARCHAR(20) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Role VARCHAR(10) CHECK (Role IN ('admin', 'customer')) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE Branch (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Location VARCHAR(255) NOT NULL,
    OpeningTime TIME NOT NULL,
    ClosingTime TIME NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_branch_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_branch_updated_at
BEFORE UPDATE ON Branch
FOR EACH ROW
EXECUTE FUNCTION update_branch_updated_at();

CREATE TABLE services (
    ID SERIAL PRIMARY KEY,
    IDBranch INT NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Type VARCHAR(50) NOT NULL,
    Detail TEXT,
    Pricing DECIMAL(10, 2) NOT NULL,
    Duration INTERVAL NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (IDBranch) REFERENCES Branch(ID)
);

CREATE OR REPLACE FUNCTION update_services_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_services_updated_at
BEFORE UPDATE ON services
FOR EACH ROW
EXECUTE FUNCTION update_services_updated_at();

CREATE TABLE Reservations (
    ID SERIAL PRIMARY KEY,
    IDUser INT NOT NULL,
    IDService INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (IDUser) REFERENCES users(ID),
    FOREIGN KEY (IDService) REFERENCES services(ID)
);

CREATE OR REPLACE FUNCTION update_reservations_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_reservations_updated_at
BEFORE UPDATE ON Reservations
FOR EACH ROW
EXECUTE FUNCTION update_reservations_updated_at();

CREATE TABLE Rating (
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

CREATE OR REPLACE FUNCTION update_rating_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_rating_updated_at
BEFORE UPDATE ON Rating
FOR EACH ROW
EXECUTE FUNCTION update_rating_updated_at();
