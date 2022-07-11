CREATE TABLE stocks (
  stock_id SERIAL PRIMARY KEY,
  name TEXT,
  price INT,
  company TEXT
);
INSERT INTO stocks(name, price, company) VALUES
('Apple', 100, 'Apple Inc.'),
('Google', 200, 'Google Inc.'),
('Microsoft', 300, 'Microsoft Inc.'),
('Amazon', 400, 'Amazon Inc.'),
('Facebook', 500, 'Facebook Inc.');
