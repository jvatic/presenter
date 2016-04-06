var Slide = function (props) {
	var slide = props.slide;
	var html = marked(slide.body);
	return (
		<article className="slide" dangerouslySetInnerHTML={{__html: html}} />
	);
};

export default Slide;
