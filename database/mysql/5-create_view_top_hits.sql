CREATE VIEW top_2_hits AS
SELECT
    (
        SELECT song_name FROM songs WHERE song_id = uh.song_id
    ) AS song,
    COUNT((SELECT song_id FROM songs WHERE song_id = uh.song_id)) AS reproductions_count
FROM users_history AS uh
GROUP BY song
ORDER BY reproductions_count DESC, song LIMIT 2;