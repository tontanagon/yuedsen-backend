ALTER TABLE poses ADD COLUMN pose_category_id BIGINT;
ALTER TABLE plans DROP COLUMN pose_category_id;
