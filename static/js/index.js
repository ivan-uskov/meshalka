function onDrinkDelete(response)
{
	if (response.success)
	{
		location.reload();
	}
	else
	{
		alert("Some error occured!");
	}
}

$(function()
{
	$(".delete_drink").click(function(){
		if (confirm("Delete this drink?"))
		{
			ajaxRequest(this.href, {}, onDrinkDelete);
		}

		return false;
	});
	
	$(".drink_container a, #add_drink_link, #add_cocktail_link").each(function(){
		$(this).fancybox({
			ajax: 
			{
				maxWidth: 1200,
				maxHeight: 800,
				frameWidth: 'auto',
				frameHeight: 600,
				autoSize	: true,
			}
		});
	});
});
