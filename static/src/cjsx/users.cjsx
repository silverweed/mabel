###
 Components for managing login and user preferences
###
window.UserWidget = React.createClass {
	getInitialState: ->
		{ user: null }

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
			<div>
			    <h3>{@state.user.name}</h3>
			    <form method='POST' action='/logout'>
			        <input type='submit' value='Log out'/>
			    </form>
			</div>
		else
			<div>
			    <LoginButton />
			</div>
}

LoginButton = React.createClass {
	getInitialState: ->
		{
			toggled: no
			regToggled: no
		}

	render: ->
		if @state.toggled
			<div className='widget'>
			    <h4 className='title'>{@title()}</h4>
			    <form method='POST' action={@action()}>
			        <input type='text'
			               name='name'
			               placeholder='Username'/>
			        <br/>
			        <input type='password'
			               name='password'
			               placeholder='Password'/>
			        <br/>
			        <RegisterInputs toggled={@state.regToggled} onClick={@showReg}/>
			        <br/>
			        <input type='submit'
			               value='Submit'/>
			        <button onClick={@toggle}>Cancel</button>
			    </form>
			</div>
		else
			<div>
			    <button onClick={@toggle}>Login</button>
			</div>

	toggle: ->
		@setState {
			toggled: !@state.toggled
			regToggled: no
		}

	showReg: -> @setState { regToggled: yes }
	title: -> if @state.regToggled then 'Register' else 'Login'
	action: -> if @state.regToggled then '/register' else '/login'
}

RegisterInputs = React.createClass {
	render: ->
		if @props.toggled
			<div>
			    <input type='password'
			           name='confirmpassword'
			           placeholder='Confirm password'/>
			    <br/>
			    <input type='text'
			           name='email'
			           placeholder='Email address'/>
			</div>
		else
			<a className='nohref' onClick={@props.onClick}>
			Need an account?
			</a>
}
