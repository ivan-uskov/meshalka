DROP TABLE IF EXISTS user;
CREATE TABLE user
(
    user_id INT UNSIGNED NOT NULL UNIQUE AUTO_INCREMENT,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(user_id)
);

DROP FUNCTION IF EXISTS `lock_user`;
DELIMITER $$
CREATE FUNCTION `lock_user`() RETURNS TINYINT(1)
DETERMINISTIC
    BEGIN
        RETURN GET_LOCK(CONCAT(DATABASE(), '.lock_user'), 60);
    END $$
DELIMITER ;

DROP FUNCTION IF EXISTS `unlock_user`;
DELIMITER $$
CREATE FUNCTION `unlock_user`() RETURNS TINYINT(1)
DETERMINISTIC
    BEGIN
        RETURN RELEASE_LOCK(CONCAT(DATABASE(), '.lock_user'));
    END $$
DELIMITER ;

DELIMITER $$
DROP FUNCTION IF EXISTS `add_user`$$
CREATE FUNCTION `add_user`(newLogin VARCHAR(255), newPassword VARCHAR(255))
    RETURNS TINYINT(1)
DETERMINISTIC
    BEGIN
        DECLARE lock_success TINYINT(1) DEFAULT 0;
        DECLARE user_not_exists TINYINT(1) DEFAULT 0;

        SELECT lock_user() INTO lock_success;

        IF (lock_success = 1) THEN
            SELECT COUNT(*) = 0 INTO user_not_exists FROM user WHERE login = newLogin;
            IF (user_not_exists = 1) THEN
                INSERT INTO user SET login = newLogin, password = newPassword;
            END IF;

            SELECT unlock_user() INTO lock_success;
        END IF;

        RETURN lock_success && user_not_exists;
    END $$

DELIMITER ;