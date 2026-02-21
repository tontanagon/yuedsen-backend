ALTER TABLE poses ADD COLUMN pose_accuracy FLOAT;
ALTER TABLE plans DROP COLUMN pose_accuracy;
