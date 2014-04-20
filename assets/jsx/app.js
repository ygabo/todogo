/** @jsx React.DOM */
var TodoList = React.createClass({
	// getInitialState: function(){
	// 	return {}
	// },
	render:function(){
		return <p>Hello World</p>
	}
})

React.renderComponent(<TodoList />, document.getElementById('app'))
