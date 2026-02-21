ALTER TABLE plans ADD COLUMN pose_accuracy FLOAT;
ALTER TABLE poses DROP COLUMN pose_accuracy;
