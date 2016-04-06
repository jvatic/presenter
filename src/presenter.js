import App from "presenter/main";

marked.setOptions({
	highlight: function (code, lang) {
		return hljs.highlight(lang, code, true).value;
	}
});

App.run(document.getElementById("main"));
