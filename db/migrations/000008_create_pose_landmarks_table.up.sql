CREATE TABLE IF NOT EXISTS pose_landmarks (
    id SERIAL PRIMARY KEY,
    pose_id INT NOT NULL,
    landmark_index INT NOT NULL,
    x FLOAT NOT NULL,
    y FLOAT NOT NULL,
    z FLOAT NOT NULL,
    visibility FLOAT,
    CONSTRAINT fk_pose_landmark
        FOREIGN KEY(pose_id) 
        REFERENCES poses(id)
        ON DELETE CASCADE
);
