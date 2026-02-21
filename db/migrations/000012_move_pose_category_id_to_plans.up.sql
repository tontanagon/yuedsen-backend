ALTER TABLE poses DROP CONSTRAINT IF EXISTS fk_poses_pose_category;
ALTER TABLE plans ADD COLUMN pose_category_id BIGINT;
ALTER TABLE poses DROP COLUMN pose_category_id;
