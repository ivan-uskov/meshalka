DROP TABLE IF EXISTS cocktail;
CREATE TABLE cocktail
(
  cocktail_id SERIAL,
  cocktail_name VARCHAR(255) NOT NULL UNIQUE,
  PRIMARY KEY(cocktail_id)
);

DROP FUNCTION IF EXISTS `lock_cocktail`;
DELIMITER $$
CREATE FUNCTION `lock_cocktail`() RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    RETURN GET_LOCK(CONCAT(DATABASE(), '.lock_cocktail'), 60);
  END $$
DELIMITER ;

DROP FUNCTION IF EXISTS `lock_cocktail`;
DELIMITER $$
CREATE FUNCTION `lock_cocktail`() RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    RETURN RELEASE_LOCK(CONCAT(DATABASE(), '.lock_cocktail'));
  END $$
DELIMITER ;

DELIMITER $$
DROP FUNCTION IF EXISTS `add_cocktail`$$
CREATE FUNCTION `add_cocktail`(newCocktailName VARCHAR(255))
  RETURNS TINYINT(1)
DETERMINISTIC
  BEGIN
    DECLARE lock_success TINYINT(1) DEFAULT 0;
    DECLARE cocktail_not_exists TINYINT(1) DEFAULT 0;

    SELECT lock_drink() INTO lock_success;
    IF (lock_success = 1) THEN
      SELECT COUNT(*) = 0 INTO cocktail_not_exists FROM cocktail WHERE cocktail_name = newCocktailName;
      IF (cocktail_not_exists = 1) THEN
        INSERT INTO cocktail SET cocktail_name = newCocktailName;
      END IF;

      SELECT unlock_drink() INTO lock_success;
    END IF;

    RETURN lock_success && cocktail_not_exists;
  END $$

DELIMITER ;