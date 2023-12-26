-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `laptops` (
                       `id` INT PRIMARY KEY AUTO_INCREMENT,
                       `cpu` VARCHAR(255) NOT NULL ,
                       `ram` INT NOT NULL ,
                       `ssd` INT NOT NULL,
                       `hdd` INT NOT NULL,
                       `graphic` INT NOT NULL,
                       `screen_size` VARCHAR(255) NOT NULL,
                       `company` VARCHAR(255) NOT NULL,
                       `price`  VARCHAR(255) NOT NULL,
                       `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       `image_url` VARCHAR(255),
                       `redirect_url` VARCHAR(255)
                       
);

-- +migrate Down
DROP TABLE `laptops`;

