DELIMITER $$

CREATE PROCEDURE find_artist_albums(IN search_name VARCHAR(50))
BEGIN
    SELECT
        (SELECT artist_name FROM artists WHERE artist_id = a.artist_id) AS artist,
        album_name AS album
    FROM albums AS a
    HAVING artist = search_name;
END $$

DELIMITER ;
