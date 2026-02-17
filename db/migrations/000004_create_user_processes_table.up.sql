CREATE TABLE IF NOT EXISTS user_processes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    pose_category_id INT NOT NULL,
    progress INT DEFAULT 1,
    status VARCHAR(50),
    total_score INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_pose_category
        FOREIGN KEY(pose_category_id) 
        REFERENCES pose_categories(id)
        ON DELETE CASCADE
);
