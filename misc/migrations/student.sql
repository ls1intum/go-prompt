CREATE TABLE public.student (
    id uuid NOT NULL,
    email character varying(100),
    first_name character varying(50),
    is_exchange_student boolean,
    last_name character varying(50),
    matriculation_number character varying(30),
    nationality character varying(10),
    public_id uuid,
    tum_id character varying(20),
    blocked_by_pm boolean,
    reason_for_blocked_by_pm oid,
    suggested_as_coach boolean,
    suggested_as_tutor boolean,
    gender character varying(255)
);