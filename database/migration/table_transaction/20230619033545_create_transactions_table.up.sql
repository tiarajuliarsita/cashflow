CREATE TABLE transactions (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  typeTransaction VARCHAR(255),
  amount INT
);


-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/miniAtm" -path database/migration up