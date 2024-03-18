INSERT INTO tours.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags) 
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'Serengeti Safari', 'Embark on an unforgettable safari through the stunning landscapes of Serengeti National Park.', 4, 59.99, 1, -11, 6, '2023-07-03 14:10:55', NULL, false, ARRAY['Adventure']);

INSERT INTO tours.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags)
VALUES ('223e4567-e89b-12d3-a456-426614174001', 'Great Barrier Reef Expedition', 'Discover the wonders beneath the ocean surface at the Great Barrier Reef.', 3, 79.99, 1, -11, 4, '2023-06-18 09:20:30', NULL, false, ARRAY['Adventure', 'Nature']);

INSERT INTO tours.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags)
VALUES ('323e4567-e89b-12d3-a456-426614174002', 'Belgrade Highlights', 'Experience the vibrant culture and rich history of Belgrade through our guided tour of its most iconic sights.', 2, 29.99, 0, -11, 3, '2023-10-18 14:30:45', NULL, false, ARRAY['Adventure', 'Sports', 'History']);

INSERT INTO tours.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags)
VALUES ('423e4567-e89b-12d3-a456-426614174003', 'Novi Sad Discovery', 'Discover the beauty and history of Novi Sad with our guided tour through its iconic landmarks.', 2, 24.99, 2, -11, 2, '2023-11-05 10:15:20', '2024-03-16 18:12:49.5261556+01:00', false, ARRAY['Adventure']);