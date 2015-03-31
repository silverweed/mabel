Greeting = React.createClass {
	render: -> return <p className="greeting">Hello, {@props.name}</p>
}

React.render <Greeting name="world"/>, document.getElementById 'content'
