function onCreateDrink(response)
{
	if (response.success)
	{
		location.reload();
	}
	else
	{
		 alert("This drink may be already exists or error occured try later!");
	}
}

$(function()
{
	$("#edit_drink_form").submit(function()
	{
		var name = $(this).find("#drink_name").val().trim();

		if (validateDrinkName(name))
		{
			ajaxRequest(this.action, {"drink_name": name}, onCreateDrink);
		}
		else
		{
			alert("This name incorrect, it should be a word, with length less than 25 charachters!");
		}

		return false;
	});
});