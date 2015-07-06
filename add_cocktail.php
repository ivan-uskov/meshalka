<?php

    define("PARAMETERS_COUNT", 3);
    define("NAME_ID", 1);
    define("DRINKS_ID", 2);
    define("LOG_PATH", "./logs/add-coctails.log");

    class CocktailInfo
    {
        private $name;
        private $drinks;

        public function __construct($name, $drinks)
        {
            $this->name = trim($name);
            $this->drinks = trim($drinks);

            if (!$this->isCocktailNameValid())
            {
                throw new Exception("Invalid cocktail name!");
            }

            if (!$this->isDrinkListValid())
            {
                throw new Exception("Invalid drinks list!");
            }
        }

        public function getName()
        {
            return $this->name;
        }

        public function getDrinkList()
        {
            return $this->drinks;
        }

        private function isCocktailNameValid()
        {
            return preg_match('/^[a-zA-Z\-\s]+$/', $this->name);
        }

        private function isDrinkListValid()
        {
            return preg_match('/^[0-9]([,][0-9])*$/', $this->drinks);
        }
    }

    class Database
    {
        private $handle;

        public function __construct()
        {
            $this->handle = new PDO("mysql:host=127.0.0.1;DBName=cocktails", "root", "1234");
            $this->handle->query("USE cocktails");
        }

        public function saveCocktailInfo(CocktailInfo $cocktailInfo)
        {
            $this->handle->beginTransaction();
            $stmt = $this->getSaveCocktailStmt($cocktailInfo);

            if (!($stmt->execute()))
            {
                $this->handle->rollBack();
                throw new Exception(var_export($stmt->errorInfo(), true));
            }

            $this->handle->commit();
        }

        private function getSaveCocktailStmt(CocktailInfo $cocktailInfo)
        {
            $name = $cocktailInfo->getName();
            $drinks = $cocktailInfo->getDrinkList();

            $stmt = $this->handle->prepare("CALL add_cocktail(:name, :drinks)");
            $stmt->bindParam(':name', $name, PDO::PARAM_STR);
            $stmt->bindParam(':drinks', $drinks, PDO::PARAM_STR);

            return $stmt;
        }
    }

    function initCocktailInfoFromArgv($values)
    {
        if (count($values) < PARAMETERS_COUNT)
        {
            throw new Exception("Invalid argument's count!");
        }

        return new CocktailInfo($values[NAME_ID], $values[DRINKS_ID]);
    }

    function saveCocktailFromArgv($argv)
    {
        try
        {
            $ci = initCocktailInfoFromArgv($argv);

            $database = new Database();
            $database->saveCocktailInfo($ci);
            echo "1";
        }
        catch (Exception $e)
        {
            file_put_contents(LOG_PATH, "[" . date("Y-m-d H:i:s") . "] Error:" . $e->getMessage() . "\n", FILE_APPEND);
            echo "0";
        }
    }

    saveCocktailFromArgv($argv);