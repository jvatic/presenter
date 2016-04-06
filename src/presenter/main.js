import History from 'marbles/history';
import MainRouter from 'presenter/routers/main';
import MainStore from 'presenter/stores/main';

var render = function ($el, component, props, children) {
	ReactDOM.render(
		React.createElement(component, props, children),
		$el
	);
};

var App = function ($el) {
	var history = this.history = new History();
	this.__handleStoreChange = this.__handleStoreChange.bind(this);
	MainStore.addChangeListener(null, this.__handleStoreChange);
	var context = {
		history: history,
		render: render.bind(null, $el)
	};
	history.register(new MainRouter({ context: context }));
};

App.run = function ($el) {
	var app = new App($el);
	app.run();
};

App.prototype.run = function () {
	this.history.start({
		context: {
			data: MainStore.getState(null)
		}
	});
	document.body.addEventListener('keypress', this.__handleKeypress.bind(this));
};

App.prototype.__handleStoreChange = function () {
	this.context = this.history.context = {
		data: MainStore.getState(null)
	};
	this.history.navigate(this.history.path, {force: true, replace: true});
};

var KEY_CODES = {
	ARROW_LEFT: 37,
	ARROW_RIGHT: 39
};

var CHAR_CODES = {
	SPACE: 32
};

App.prototype.__handleKeypress = function (e) {
	var currentSlideIndex = this.context.data.currentSlideIndex;
	var lastSlideIndex = this.context.data.slides.length-1;
	var nextSlideIndex = Math.min(currentSlideIndex + 1, lastSlideIndex);
	var prevSlideIndex = Math.max(currentSlideIndex - 1, 0);
	if (e.keyCode === KEY_CODES.ARROW_LEFT) {
		e.preventDefault();
		this.history.navigate('/slides/'+ String(prevSlideIndex));
	} else if (e.keyCode === KEY_CODES.ARROW_RIGHT || e.charCode === CHAR_CODES.SPACE) {
		e.preventDefault();
		this.history.navigate('/slides/'+ String(nextSlideIndex));
	}
};

export default App;
