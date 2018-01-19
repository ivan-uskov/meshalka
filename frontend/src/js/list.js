"use strict";

const DRINK_TYPES = {
	DRINK: 0,
	COCKTAIL: 1,
};

function requestErrorHandler(error) {
	if (error.status !== 401) {
		throw error;
	}
	window.location.href = "/";
}

function initList(drinks) {
	drinks = drinks["drinks"];
	$("#drinks").html("");
	for (const drink of drinks) {
		createDrinkCard(drink, drinks);
	}
}

function createDrinkCard(drink, drinks) {
	let card = $("#drink-template").html();
	card = card.replace(/#id#/g, drink["id"]);
	card = card.replace(/#description#/g, drink["name"]);
	card = $(card);
	$("#drinks").append(card);

	card.find(".delete-button").click(() => {
		removeDrink(card, drink["id"]);
	});

	const input = card.find(`#${drink.id}`);
	const editField = card.find(".mdl-card__supporting-text");
	editField.attr("disabled", true);
	card.find(".edit-button").click(() => {
		editField.removeAttr("disabled");
		input.focus();
	});
	/*input.focusout(() => {
		editField.attr("disabled", true);
		editDrinkName(drink["id"], input.val());
	});*/
	input.on("input", () => {
		card.find(".mdl-card__title-text").text(input.val());
	});
	if (drink["type"] === DRINK_TYPES.COCKTAIL) {
		addComponentsEditGroup(editField, drink, drinks);
	}
}

function addComponentsEditGroup(editField, cocktail, drinks) {
	const componentsHolder = $("<div class=`components-holder`></div>");
	const addComponentButton = $($("#add-component-button-template").html());
	editField.append(componentsHolder);
	editField.append(addComponentButton);
	addComponentButton.on('click', () => {
		cocktail["components"].push(getDefaultDrink(cocktail, drinks));
		editCockteil(cocktail);
	});
	window.addEventListener("drinks-change", () => {
        invalidateComponents(componentsHolder, cocktail, drinks);
	});
	invalidateComponents(componentsHolder, cocktail, drinks);
	componentsHolder.on("cocktail-changed", () => {
		editCockteil(cocktail);
	});
}

function invalidateComponents(componentsHolder, drink, drinks) {
	componentsHolder.html("");
	const componentsModels = drink["components"];
	const checkedIds = [];
	const components = [];
	for (let i = 0; i < componentsModels.length; ++i) {
		const componentModel = componentsModels[i];
		checkedIds.push(componentModel["id"]);
		const component = createComponent(componentModel, drinks);
		componentsHolder.append(component);
		component.on("clear", () => {
			component.remove();
			componentsModels.splice(i, 1);
			componentsHolder.trigger("cocktail-changed");
		});
		component.on("selected-item-change", () => componentsHolder.trigger("cocktail-changed"));
		component.on("volume-change", () => componentsHolder.trigger("cocktail-changed"));
		component.on("selection-change", () => componentsHolder.trigger("cocktail-changed"));
		components.push(component);
	}
	for (const component of components)
	{
		for (const option of component.find("select").children())
		{
			if (!option.attr("checked") && checkedIds.indexOf(option.val()) !== -1)
			{
				option.remove();
			}
		}
	}
}

function createComponent(componentModel, drinks) {
	let component = $("#component-template").html();
	const id = Math.floor(Math.random() * 10000);
	component = $(component.replace(/#id#/g, id));
	const select = component.find("component-select");
	const optionsTemplate = "<option value='#id#'>#description#</option>";
	for (const drink of drinks) {
		if (drinks["type"] === DRINK_TYPES.DRINK) {
			const option = $(optionsTemplate.replace(/#id#/g, drink["id"]));
			select.append(option);
			option.attr("checked", drink["id"] === componentModel[id]);
		}
	}
	component.find(".clear-button").click(() => {
		component.trigger("clear");
	});
	const input = component.find(".mdl-textfield__input");
	input.on("input", () => {
		componentModel["volume"].text(input.val());
	});
	input.focusout(() => {
		component.trigger("volume-change");
	});
	select.change(() => {
		component.trigger("selection-change");
		componentModel["id"] = select.val();
	});
	component.find(".mdl-textfield__input").val(componentModel["volume"]);
	return component;
}

function getDefaultDrink(cocktail, drinks) {
	let filteredDrinks = drinks.filter((drink) => drink.type === DRINK_TYPES.DRINK);

	let drink = filteredDrinks[Math.floor(Math.random() * filteredDrinks.length)];
    return {id: drink["id"], volume: 1};
}

function editCockteil(drink) {
	console.log(drink);
	const components = drink["components"];
	const restructuredComponents = {};
	for (let i = 0; i < components.length; ++i) {
		restructuredComponents[components[i]["id"]] = +components[i]["volume"];
	}
	drink["components"] = JSON.stringify(restructuredComponents);
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data: JSON.stringify({
			action: "edit_cocktail_components",
			data: JSON.stringify(drink)
		}),
		success: () => {},
		error: requestErrorHandler
	});
}

function editDrinkName(id, name) {
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data: JSON.stringify({
			action: "edit_drink",
			data: JSON.stringify({id, name: name})
		}),
		success: window.dispatchEvent(new CustomEvent("drinks-change")),
		error: requestErrorHandler
	});
}

function removeDrink(card, id) {
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data: JSON.stringify({
			action: "remove_drink",
			data: `${id}`
		}),
		success: () => card.remove(),
		error: requestErrorHandler,
	});
}

function addDrink() {
	const name = $("#default-drink-name").val();
	const type = +$("input[name=drink-types]:checked").val();
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data: JSON.stringify({
			action: "add_drink",
			data: JSON.stringify({name, type})
		}),
		success: invalidateList,
		error: requestErrorHandler
	});
}

function invalidateList() {
	$.ajax({
		type: "POST",
		url: "/api",
		dataType: "json",
		data: JSON.stringify({action: "list"}),
		success: initList,
		error: requestErrorHandler
	});
}

invalidateList();
$("#add-drink").click(addDrink);