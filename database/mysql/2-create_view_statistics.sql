CREATE VIEW music_statistics AS
SELECT COUNT(song_name) AS songs,
    (SELECT COUNT(artist_name) FROM artists) AS artists,
    (SELECT COUNT(album_name) FROM albums) AS albums
FROM songs;
