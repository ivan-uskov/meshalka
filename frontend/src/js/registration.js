"use strict";
function showDrinkList() {
	window.location.href = "list";
}

function registrationErrorHandler(error) {
	if (error.status !== 400) {
		throw error;
	}
	showWrongDataMessage();
}

function showWrongDataMessage() {
	const messageBox = document.querySelector("#message-toast");
	var data = {
		message: "There is user with same login. Create another login.",
	};
	messageBox.MaterialSnackbar.showSnackbar(data);
}

$("#register-profile").click(() => {
	const login = $("#username").val();
	const pass = $("#userpass").val();
	const data = JSON.stringify({
		action: "register",
		data: JSON.stringify({login, pass})
	});
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data,
		success: showDrinkList,
		error: registrationErrorHandler
	});
});
