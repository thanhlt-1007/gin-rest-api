-- Create "tokens" table
CREATE TABLE `tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `tokenable_id` bigint unsigned NULL,
  `tokenable_type` varchar(255) NULL,
  `access_token` varchar(255) NULL,
  `refresh_token` varchar(255) NULL,
  `expires_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_tokens_access_token` (`access_token`),
  UNIQUE INDEX `idx_tokens_refresh_token` (`refresh_token`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `email` varchar(255) NULL,
  `name` varchar(255) NULL,
  `encrypted_password` varchar(255) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_deleted_at` (`deleted_at`),
  UNIQUE INDEX `idx_users_email` (`email`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
