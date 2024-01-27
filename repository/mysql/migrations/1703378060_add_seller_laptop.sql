-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `seller_laptop` (
                       `id` INT PRIMARY KEY AUTO_INCREMENT,
                       `seller_ref` INT ,
                       `laptop_ref` INT,
                       FOREIGN KEY(`seller_ref`)REFERENCES users(id),
                       FOREIGN KEY(`laptop_ref`) REFERENCES laptops(id)
                       
);

-- +migrate Down
DROP TABLE `seller_laptop`;

