ALTER TABLE poses ALTER COLUMN pose_image TYPE BYTEA USING pose_image::bytea;
