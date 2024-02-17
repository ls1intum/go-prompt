CREATE TABLE public.developer_application (
    id uuid NOT NULL,
    current_semester smallint,
    experience oid,
    motivation oid,
    application_assessment_id uuid,
    course_iteration_id uuid,
    project_team_id uuid,
    post_kickoff_submission_id uuid,
    study_program character varying(255),
    study_degree character varying(255),
    student_id uuid NOT NULL,
    devices public.device[],
    courses_taken public.course[],
    english_language_proficiency character varying(255),
    german_language_proficiency character varying(255),
    CONSTRAINT developer_application_current_semester_check CHECK (((current_semester <= 99) AND (current_semester >= 1)))
);