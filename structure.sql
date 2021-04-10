CREATE TABLE zone1.feedback (
    id uuid NOT NULL,
    rating int2 NOT NULL,
    diagnosis_understood bool NOT NULL,
    explanation_improvement text NULL,
    feelings text NOT NULL,
    time_reported timestamp(0) NOT NULL,
    diagnosis varchar(50) NOT NULL,
    CONSTRAINT feedback_pkey PRIMARY KEY (id)
);
