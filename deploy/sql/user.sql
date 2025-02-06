create table `users`(
                        `id` varchar(24) collate utf8mb4_unicode_ci not null,
                        `avatar` varchar(191) collate utf8mb4_unicode_ci not null default '',
                        `nickname` varchar(24) collate utf8mb4_unicode_ci not null,
                        `phone` varchar(20) collate utf8mb4_unicode_ci not null,
                        `password` varchar(191) collate utf8mb4_unicode_ci DEFAULT null,
                        `status` tinyint collate utf8mb4_unicode_ci default null,
                        `gender` tinyint collate utf8mb4_unicode_ci default null,
                        `created_at` timestamp null default null,
                        `updated_at` timestamp null default null,
                        primary key (`id`)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;
