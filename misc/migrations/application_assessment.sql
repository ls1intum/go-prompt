CREATE TABLE public.application_assessment (
    id uuid NOT NULL,
    assessed boolean,
    assessment_score integer,
    technical_challenge_programming_score double precision,
    technical_challenge_quiz_score double precision,
    status character varying(255) DEFAULT 'NOT_ASSESSED'::public.application_status NOT NULL
);