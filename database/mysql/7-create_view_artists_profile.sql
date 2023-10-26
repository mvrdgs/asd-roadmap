CREATE VIEW artists_profile AS
SELECT
    (SELECT artist_name FROM artists WHERE artist_id = a.artist_id) AS artist,
    album_name AS album,
    (SELECT COUNT(*) FROM followers WHERE artist_id = a.artist_id) AS followers
FROM albums AS a
ORDER BY followers DESC, artist, album;
