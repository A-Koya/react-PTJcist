USE cist_ptj;
CREATE TABLE IF NOT EXISTS `users` 
(
    `id` VARCHAR(8) PRIMARY KEY,
    `password` VARCHAR(32) NOT NULL
);