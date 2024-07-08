SELECT 
    ui.id AS user_id,
    ui.email,
    ui.first_name,
    ui.surname,
    ui.mobile_phone,
    up.rating,
    up.description AS profile_description,
    up.done_project,
    up.given_project
FROM 
    user_identity ui
JOIN 
    user_profile up ON ui.id = up.user_id
WHERE 
    ui.id = 2;