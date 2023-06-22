CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50),
  email VARCHAR(255),
  pin int (255),
  accountNumber int(255),
  phoneNumber int(255),
  balance int(255),
  born_date timestamp
);

-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/miniAtm" -path database/migration up


-- migrate create -ext sql -dir migrations create_transactions_table