CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO zone1.feedback (id, rating, diagnosis_understood, feelings, diagnosis, time_reported)
VALUES
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E10-E14.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'I10', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days')),
(uuid_generate_v4(), random() * 10, random() > 0.5, 'Lorem ipsum dolor sit amet', 'E55.9', now()-random()*(interval '1 days'))
