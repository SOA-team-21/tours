DELETE FROM public.required_times;
DELETE FROM public.key_points;
DELETE FROM public.tours; 

INSERT INTO public.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags) 
VALUES (-1, 'Serengeti Safari', 'Embark on an unforgettable safari through the stunning landscapes of Serengeti National Park.', 4, 59.99, 1, -11, 6, '2023-07-03 14:10:55', NULL, false, ARRAY['Adventure']);

INSERT INTO public.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags)
VALUES (-2, 'Great Barrier Reef Expedition', 'Discover the wonders beneath the ocean surface at the Great Barrier Reef.', 3, 79.99, 1, -11, 4, '2023-06-18 09:20:30', NULL, false, ARRAY['Adventure', 'Nature']);

INSERT INTO public.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags)
VALUES (-3, 'Belgrade Highlights', 'Experience the vibrant culture and rich history of Belgrade through our guided tour of its most iconic sights.', 2, 29.99, 0, -11, 3, '2023-10-18 14:30:45', NULL, false, ARRAY['Adventure', 'Sports', 'History']);

INSERT INTO public.tours(id, name, description, difficult, price, status, author_id, length, publish_time, archive_time, my_own, tags)
VALUES (-4, 'Novi Sad Discovery', 'Discover the beauty and history of Novi Sad with our guided tour through its iconic landmarks.', 2, 24.99, 2, -11, 2, '2023-11-05 10:15:20', '2024-03-16 18:12:49.5261556+01:00', false, ARRAY['Adventure']);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-1, -2, -16.3159, 145.8737, 'Coral Gardens', 'Dive into the vibrant Coral Gardens and witness the beauty of the Great Barrier Reef.', 'https://www.jordanrobins.com.au/wp-content/uploads/2020/10/Coral-Gardens-_-Lord-Howe-Island-V1.1.jpg', false);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-2, -2, -16.3159, 145.8737, 'Marine Life Paradise', 'Explore the rich marine life paradise within the Great Barrier Reef.', 'https://i0.wp.com/post.medicalnewstoday.com/wp-content/uploads/sites/3/2020/04/iStock-1170804921-1024x683.jpg?w=1155&h=2969', false);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-3, -3, 44.8231, 20.4484, 'Kalemegdan Fortress', 'Begin your journey at the historic Kalemegdan Fortress, offering panoramic views of the confluence of the Sava and Danube rivers.', 'https://belgrade-beat.rs/photos/areas/11/c-1524480090.jpg', false);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-4, -3, 44.8237, 20.4615, 'Skadarlija District', 'Immerse yourself in the bohemian atmosphere of Skadarlija, famous for its cobblestone streets and traditional Serbian restaurants.', 'https://www.naturala.hr/wp-content/uploads/2019/04/beograd-skadarlija-800x533.jpg', false);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-5, -3, 44.8154, 20.4604, 'Republic Square', 'Visit the bustling Republic Square, surrounded by important landmarks such as the National Museum and the National Theatre.', 'https://www.011info.com/images/up/Trg_Republike_i_spomenik_Knezu_Mihailu_(4)_-_Pavle_Kaplanec_copy.jpg', false);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-6, -4, 45.2516, 19.8682, 'Petrovaradin Fortress', 'Explore the historic Petrovaradin Fortress, known for its stunning views of the Danube River.', 'https://www.exitfest.org/wp-content/uploads/2020/07/DJI_0093-1.jpg', false);

INSERT INTO public.key_points(id, tour_id, latitude, longitude, name, description, picture, public)
VALUES (-7, -4, 45.2615, 19.8328, 'Dunavska Street', 'Stroll through the charming Dunavska Street, filled with cafes, shops, and vibrant street art.', 'https://live.staticflickr.com/65535/51147965485_6e583bbc62_h.jpg', false);

INSERT INTO public.required_times(id, tour_id, transport, minutes)
VALUES (-1, -1, 1, 15);

INSERT INTO public.required_times(id, tour_id, transport, minutes)
VALUES (-2, -1, 1, 15);

INSERT INTO public.required_times(id, tour_id, transport, minutes)
VALUES (-3, -1, 1, 15);

INSERT INTO public.required_times(id, tour_id, transport, minutes)
VALUES (-4, -1, 2, 20);