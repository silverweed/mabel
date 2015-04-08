###
 Components for managing login and user preferences
###
window.UserWidget = React.createClass {
	getInitialState: ->
		return {
			user: null
		}

	componentDidMount: ->
		# Retreive user data from the server
		qwest.post(@props.source).then ((resp) ->
			if @isMounted()
				data = JSON.parse resp
				@setState {
					user: data
				}
		).bind this

	render: ->
		if @state.user?.status.authenticated
			return (
				<div>
				    <h3>{@state.user.name}</h3>
				    <form method='POST' action='/logout'>
				        <input type='submit' value='Log out'/>
				    </form>
				</div>
			)
		else
			return (
				<div>
				    <LoginButton />
				</div>
			)
}

LoginButton = React.createClass {
	getInitialState: ->
		return { toggled: false }

	render: ->
		if @state.toggled
			return (
				<div>
				    <form method='POST' action='/login'>
				        <input type='text'
				               name='name'
				               placeholder='Username'/>
				        <br/>
				        <input type='password'
				               name='password'
				               placeholder='Password'/>
				        <br/>
				        <input type='submit'
				               value='Submit'/>
				        <button onClick={@toggle}>Cancel</button>
				    </form>
				</div>
			)
		else
			return (
				<div>
				    <button onClick={@toggle}>Login</button>
				</div>
			)

	toggle: -> @setState({ toggled: !@state.toggled })
}
