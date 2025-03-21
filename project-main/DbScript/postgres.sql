DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
  customer_id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  date_of_birth DATE NOT NULL,
  city VARCHAR(100) NOT NULL,
  zipcode VARCHAR(10) NOT NULL,
  status BOOLEAN NOT NULL DEFAULT TRUE
);
INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES
  (2000, 'Steve', '1978-12-15', 'Delhi', '110075', TRUE),
  (2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', TRUE),
  (2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', TRUE),
  (2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', FALSE),
  (2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', TRUE),
  (2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', FALSE);

DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
  account_id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  opening_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  account_type VARCHAR(10) NOT NULL,
  amount DECIMAL(10,2) NOT NULL,
  status BOOLEAN NOT NULL DEFAULT TRUE,
  FOREIGN KEY (customer_id) REFERENCES customers (customer_id)
);
INSERT INTO accounts (account_id, customer_id, opening_date, account_type, amount, status) VALUES
  (95470, 2000, '2020-08-22 10:20:06', 'saving', 6823.23, TRUE),
  (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, TRUE),
  (95472, 2001, '2020-08-09 10:35:22', 'saving', 7000, TRUE),
  (95473, 2001, '2020-08-09 10:38:22', 'saving', 5861.86, TRUE);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
  transaction_id SERIAL PRIMARY KEY,
  account_id INT NOT NULL,
  amount DECIMAL(10,2) NOT NULL,
  transaction_type VARCHAR(10) NOT NULL,
  transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  username VARCHAR(20) PRIMARY KEY,
  password VARCHAR(20) NOT NULL,
  role VARCHAR(20) NOT NULL,
  customer_id INT,
  created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers (customer_id)
);
INSERT INTO users (username, password, role, customer_id, created_on) VALUES
  ('admin', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),
  ('2001', 'abc123', 'user', 2001, '2020-08-09 10:27:22'),
  ('2000', 'abc123', 'user', 2000, '2020-08-09 10:27:22');

DROP TABLE IF EXISTS refresh_token_store;
CREATE TABLE refresh_token_store (
  refresh_token VARCHAR(300) PRIMARY KEY,
  created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);