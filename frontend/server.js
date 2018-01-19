"use strict";
const express = require("express");
const fs = require("fs");
const router = express();

router.get("/", function(req, res, next) {
	res.sendFile(__dirname + "/src/logination.html");
});

router.get("/registration", function(req, res, next) {
	res.sendFile(__dirname + "/src/registration.html");
});

router.get("/list", function(req, res, next) {
	res.sendFile(__dirname + "/src/list.html");
});

loadDirectory("css");
loadDirectory("js");

function loadDirectory(dir) {
	fs.readdir(`${__dirname}/src/${dir}`, (err, items) => {
		for (const item of items) {
			const relativeItemPath = `/${dir}/${item}`;
			console.log(relativeItemPath);
			router.get(relativeItemPath, function(req, res, next) {
				res.sendFile(`${__dirname}/src/${relativeItemPath}`);
			});
		}
	});
}

router.listen(3000);