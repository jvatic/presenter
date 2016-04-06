import Slide from 'presenter/views/slide';

var Slides = function (props) {
	return (
		<section className="slides">
			{props.data.slides.length > 0 ? (
				<Slide slide={props.data.slides[props.data.currentSlideIndex]} />
			) : null}
		</section>
	);
};

export default Slides;
