function ajaxRequest(url, data, handler)
{
	$.post(url, data).done(function(response){
		handler(JSON.parse(response));
	});
}