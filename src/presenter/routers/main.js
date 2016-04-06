import Router from 'marbles/router';
import Dispatcher from 'marbles/dispatcher';
import SlidesComponent from 'presenter/views/slides';

var MainRouter = Router.createClass({
	routes: [
		{ path: '/', handler: 'root' },
		{ path: '/slides/:idx', handler: 'root' }
	],

	root: function (params, opts, ctx) {
		var index = parseInt(String(params[0].idx), 10) || 0;
		if (index !== ctx.data.currentSlideIndex) {
			Dispatcher.dispatch({
				name: 'SET_CURRENT_SLIDE_INDEX',
				index: index
			});
		} else if ( !params[0].idx ) {
			this.history.navigate("/slides/"+ String(index), {replace: true});
			return;
		}
		this.context.render(SlidesComponent, { data: ctx.data });
	}
});

export default MainRouter;
