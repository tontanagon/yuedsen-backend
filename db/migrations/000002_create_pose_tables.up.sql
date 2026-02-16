CREATE TABLE IF NOT EXISTS pose_categories (
    id SERIAL PRIMARY KEY,
    category VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS poses (
    id SERIAL PRIMARY KEY,
    pose_name VARCHAR(255) NOT NULL,
    pose_image BYTEA,
    pose_description TEXT,
    pose_condition TEXT,
    pose_category_id INT NOT NULL,
    pose_accuracy FLOAT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_pose_category
        FOREIGN KEY(pose_category_id) 
        REFERENCES pose_categories(id)
        ON DELETE CASCADE
);
