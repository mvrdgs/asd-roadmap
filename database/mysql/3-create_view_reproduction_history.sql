CREATE VIEW user_reproduction_history AS
SELECT (
    SELECT username FROM users WHERE user_id = us.user_id
) AS user,
    (
        SELECT song_name FROM songs WHERE song_id = us.song_id
    ) AS song
FROM users_history AS us
ORDER BY user, song;