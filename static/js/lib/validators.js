function validateDrinkName(name)
{
	return /^[a-zA-Z0-9-\s]+$/.test(name) && name.length <= 25;
}

function validateCoctailName(name)
{
	return validateDrinkName(name);
}