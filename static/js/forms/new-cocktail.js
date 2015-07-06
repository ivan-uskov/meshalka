function onCreateCocktail(response)
{
	if (response.success)
	{
		location.reload();
	}
	else
	{
		 alert("This cocktail may be already exists or error occured try later!");
	}
}

function getSelectedDrinks(form)
{
	var selected = "";
	$(form).find(".drink_checkbox").each(function(){
		if (this.checked)
		{
			selected += this.name + ",";
		}
	});

	if (selected.length > 0)
	{
		selected = selected.slice(0, -1);
	}

	return selected;
}

$(function()
{
	$("#new_cocktail_form").submit(function()
	{
		var name = $(this).find("#cocktail_name").val().trim();
		var drinks = getSelectedDrinks(this);

		if (validateCoctailName(name) && drinks.length > 0)
		{
			var data = 
			{
				cocktail_name: name,
				drinks: drinks
			};

			ajaxRequest(this.action, data, onCreateCocktail);
		}
		else
		{
			alert("This name incorrect or drink hasn't selected!");
		}

		return false;
	});
});