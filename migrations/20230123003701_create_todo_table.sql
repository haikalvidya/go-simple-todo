-- migrate:up
CREATE TABLE `todos` ( 
    `id` INT NOT NULL AUTO_INCREMENT,
    `title` TEXT NULL,
    `activity_group_id` INT NOT NULL,
    `is_active` TINYINT NULL DEFAULT 1,
    `priority` TEXT NULL DEFAULT 'very-high',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;