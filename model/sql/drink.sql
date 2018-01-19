DROP TABLE IF EXISTS drink;
CREATE TABLE drink
(
  drink_id SERIAL,
  drink_name VARCHAR(255) NOT NULL UNIQUE,
  PRIMARY KEY(drink_id)
);

DROP FUNCTION IF EXISTS `lock_drink`;
DELIMITER $$
CREATE FUNCTION `lock_drink`() RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    RETURN GET_LOCK(CONCAT(DATABASE(), '.lock_drink'), 60);
  END $$
DELIMITER ;

DROP FUNCTION IF EXISTS `unlock_drink`;
DELIMITER $$
CREATE FUNCTION `unlock_drink`() RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    RETURN RELEASE_LOCK(CONCAT(DATABASE(), '.lock_drink'));
  END $$
DELIMITER ;

DELIMITER $$
DROP FUNCTION IF EXISTS `add_drink`$$
CREATE FUNCTION `add_drink`(newDrinkName VARCHAR(255))
  RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    DECLARE lock_success TINYINT(1) DEFAULT 0;
    DECLARE drink_not_exists TINYINT(1) DEFAULT 0;

    SELECT lock_drink() INTO lock_success;

    IF (lock_success = 1) THEN
      SELECT COUNT(*) = 0 INTO drink_not_exists FROM drink WHERE drink_name = newDrinkName;
      IF (drink_not_exists = 1) THEN
        INSERT INTO drink SET drink_name = newDrinkName;
      END IF;

      SELECT unlock_drink() INTO lock_success;
    END IF;

    RETURN lock_success && drink_not_exists;
  END $$

DELIMITER ;

DELIMITER $$
DROP FUNCTION IF EXISTS `remove_drink`$$
CREATE FUNCTION `remove_drink`(drinkId BIGINT UNSIGNED)
  RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    DECLARE lock_success TINYINT(1) DEFAULT 0;
    DECLARE drink_exists TINYINT(1) DEFAULT 0;

    SELECT lock_drink() INTO lock_success;

    IF (lock_success = 1) THEN
      SELECT COUNT(*) > 0 INTO drink_exists FROM drink WHERE drink_id = drinkId;
      IF (drink_exists = 1) THEN
        DELETE FROM drink WHERE drink_id = drinkId;
      END IF;

      SELECT unlock_drink() INTO lock_success;
    END IF;

    RETURN lock_success && drink_exists;
  END $$

DELIMITER ;

DELIMITER $$
DROP FUNCTION IF EXISTS `edit_drink`$$
CREATE FUNCTION `edit_drink`(drinkId BIGINT UNSIGNED, newDrinkName VARCHAR(255))
  RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    DECLARE lock_success TINYINT(1) DEFAULT 0;
    DECLARE drink_exists TINYINT(1) DEFAULT 0;
    DECLARE name_busy TINYINT(1) DEFAULT 0;

    SELECT lock_drink() INTO lock_success;
    IF (lock_success = 1) THEN
      SELECT COUNT(*) > 0 INTO name_busy FROM drink WHERE drink_id <> drinkId AND drink_name = newDrinkName;
      SELECT COUNT(*) > 0 INTO drink_exists FROM drink WHERE drink_id = drinkId;
      IF (NOT name_busy AND drink_exists) THEN
        UPDATE drink SET drink_name = newDrinkName WHERE drink_id = drinkId;
      END IF;
      SELECT unlock_drink() INTO lock_success;
    END IF;

    RETURN lock_success && drink_exists && NOT name_busy;
  END $$

DELIMITER ;