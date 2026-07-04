CREATE TABLE comments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    post_id BIGINT UNSIGNED NOT NULL,
    author_id BIGINT UNSIGNED,
    author_name VARCHAR(255),
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    KEY idx_comments_post_id (post_id),
    KEY idx_comments_author_id (author_id),
    CONSTRAINT fk_comments_post_id FOREIGN KEY (post_id) REFERENCES posts (id),
    CONSTRAINT fk_comments_author_id FOREIGN KEY (author_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
