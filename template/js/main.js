var sections = document.getElementsByTagName('section').item(0);
var uls = sections.getElementsByTagName('ul').item(0);
console.log(uls);
var lis = uls.getElementsByTagName('li');
var labels = lis[0].getElementsByTagName('label');
labels[0].onclick = function (event) {
	if (event && event.stopPropagation) {
		console.log(this);
		event.stopPropagation();
	} else {
		window.event.cancelBubble = true;
	}
}