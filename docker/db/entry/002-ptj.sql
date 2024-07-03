USE cist_ptj;
CREATE TABLE IF NOT EXISTS `part_time_job` 
(
    `id` CHAR(36) PRIMARY KEY,
    `date` DATE,
    `start_time` TIME,
    `finish_time` TIME,
    `break_time` TIME,
    `working_time` TIME,
    `subject` TIME,
    `duty` TEXT
);