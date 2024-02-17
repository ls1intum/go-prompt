-- name: ListDeveloperApplication :many
WITH student_a AS (
    SELECT 
    s.id,
    convert_from(lo_get(s.reason_for_blocked_by_pm), 'UTF-8') as reason_for_blocked_by_pm
FROM student s)
SELECT
    d.id,
    current_semester,
    convert_from(lo_get(experience), 'UTF-8') as experience,
    convert_from(lo_get(motivation), 'UTF-8') as motivation,
    course_iteration_id,
    project_team_id,
    post_kickoff_submission_id,
    study_program,
    study_degree,
    devices,
    courses_taken,
    english_language_proficiency,
    german_language_proficiency,
    sqlc.embed(s),
    sqlc.embed(a)
FROM developer_application d
LEFT JOIN student_a s ON s.id = d.student_id
LEFT JOIN application_assessment a on a.id = d.application_assessment_id;

