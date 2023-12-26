-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `user_laptop` (
                       `id` INT PRIMARY KEY AUTO_INCREMENT,
                       `user_ref` INT ,
                       `laptop_ref` INT,
                       FOREIGN KEY(`user_ref`)REFERENCES users(id),
                       FOREIGN KEY(`laptop_ref`) REFERENCES laptops(id)
                       
);

-- +migrate Down
DROP TABLE `user_laptop`;

