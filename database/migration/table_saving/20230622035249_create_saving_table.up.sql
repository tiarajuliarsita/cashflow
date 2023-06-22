CREATE TABLE saving (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  amount INT(255),
  date TIMESTAMP,
 user_id INT(255),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/miniAtm" -path database/migration up