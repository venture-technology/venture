ALTER TABLE drivers
ADD COLUMN profile_image TEXT DEFAULT 'null';

ALTER TABLE responsible
ADD COLUMN profile_image TEXT DEFAULT 'null';

ALTER TABLE children
ADD COLUMN profile_image TEXT DEFAULT 'null';

ALTER TABLE schools
ADD COLUMN profile_image TEXT DEFAULT 'null';