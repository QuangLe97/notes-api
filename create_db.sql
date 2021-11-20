CREATE TABLE `notes`
(
    `id`           INTEGER PRIMARY KEY AUTOINCREMENT,
    `name`         VARCHAR(255)  NOT NULL,
    `content`      VARCHAR(1000) NOT NULL,
    `time_created` DATETIME      NOT NULL,
    `time_updated` DATETIME      NOT NULL
);