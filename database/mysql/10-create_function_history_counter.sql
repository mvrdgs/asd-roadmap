DELIMITER $$

CREATE FUNCTION music_history_count(user_id INT)
    RETURNS INT READS SQL DATA
BEGIN
    DECLARE total_count INT;
    SELECT COUNT(*) AS music_history_count FROM users_history uh WHERE uh.user_id = user_id INTO total_count;
    RETURN total_count;
END $$

DELIMITER ;

SELECT music_history_count(3);
