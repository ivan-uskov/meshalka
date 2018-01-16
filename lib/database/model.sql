CREATE DATABASE IF NOT EXISTS cocktails;

USE cocktails

CREATE TABLE IF NOT EXISTS drink 
(
	drink_id SERIAL,
	drink_name VARCHAR(255) NOT NULL UNIQUE,
	create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(drink_id)
);

CREATE TABLE IF NOT EXISTS cocktail
(
	cocktail_id SERIAL,
	cocktail_name VARCHAR(255) NOT NULL UNIQUE,
	create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(cocktail_id)
);

CREATE TABLE IF NOT EXISTS cocktail_element
(
	element_id SERIAL,
	cocktail_id BIGINT UNSIGNED NOT NULL,
	drink_id BIGINT UNSIGNED NOT NULL,
	FOREIGN KEY(cocktail_id) REFERENCES cocktail(cocktail_id) ON DELETE CASCADE,
	FOREIGN KEY(drink_id) REFERENCES drink(drink_id) ON DELETE CASCADE,
	PRIMARY KEY(element_id)
);

DELIMITER $$
DROP PROCEDURE IF EXISTS add_cocktail$$

CREATE PROCEDURE add_cocktail(cocktailName VARCHAR(255), drinkIdString VARCHAR(500))
BEGIN
  INSERT INTO cocktail
  SET
    cocktail_name = cocktailName;

  SET @drinks = CONCAT(drinkIdString, ',');
  SET @qry = CONCAT('INSERT INTO cocktail_element SET cocktail_id = ', LAST_INSERT_ID(), ', drink_id = ?');
  PREPARE stmt FROM @qry;

  WHILE STRCMP(@drinks, '') DO
    SET @drink_id = SUBSTRING_INDEX(@drinks, ',', 1);
    SET @drinks = REPLACE(@drinks, CONCAT(@drink_id, ','), '');
    EXECUTE stmt USING @drink_id;
  END WHILE;

  DEALLOCATE PREPARE stmt;
END $$

DELIMITER ;

