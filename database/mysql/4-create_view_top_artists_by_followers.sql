CREATE VIEW top_3_artists AS
SELECT (
    SELECT artist_name FROM artists WHERE artist_id = f.artist_id
) AS artist,
    COUNT((SELECT username FROM users WHERE user_id = f.user_id)) AS followers
FROM followers AS f
GROUP BY artist
ORDER BY followers DESC, artist LIMIT 3;