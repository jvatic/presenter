import Store from 'marbles/store';
import Dispatcher from 'marbles/dispatcher';
import Client from 'presenter/client';

var MainStore = Store.createClass({
	getInitialState: function () {
		return {
			slides: [],
			currentSlideIndex: 0
		};
	},

	didInitialize: function () {
		Client.getSlides().then(function (data) {
			this.setState({ slides: data.slides || [] });
		}.bind(this));
	},

	handleEvent: function (event) {
		switch (event.name) {
		case 'SET_CURRENT_SLIDE_INDEX':
			this.setState({
				currentSlideIndex: event.index
			});
			break;
		}
	}
});

MainStore.registerWithDispatcher(Dispatcher);

export default MainStore;
