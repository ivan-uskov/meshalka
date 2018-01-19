"use strict";
function showDrinkList() {
	window.location.href = "list";
}

function loginationErrorHandler(error) {
	if (error.status !== 401) {
		throw error;
	}
	showWrongDataMessage();
}

function showWrongDataMessage() {
	const messageBox = document.querySelector("#message-toast");
	var data = {
		message: "Wrong login or password.",
	};
	messageBox.MaterialSnackbar.showSnackbar(data);
}

$("#log-in").click(() => {
	const login = $("#username").val();
	const pass = $("#userpass").val();
	const data = JSON.stringify({
		action: "login",
		data: JSON.stringify({login, pass})
	});
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data,
		success: showDrinkList,
		error: loginationErrorHandler
	});
});