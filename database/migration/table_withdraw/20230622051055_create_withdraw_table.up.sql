CREATE TABLE withdraw (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  amount INT(255),
  date TIMESTAMP,
 user_id INT(255),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
