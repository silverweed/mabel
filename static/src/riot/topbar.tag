<topbar>
	<div id='topbar'>
		<user-widget source='/userdata'/>
		<search-bar/>
	</div>
</topbar>

<search-bar>
	<form class='searchBar' method='POST' source='/search'>
		<input type='text' placeholder='Search by tag' />
		<button type='submit'>Search</button>
	</form>
</search-bar>

<user-widget>
	<div class='right'>
		<div if={authenticated} class='right'>
			<h3>{user.name}</h3>
			<form method='POST' action='/logout'>
				<button type='submit'>Log out</button>
			</form>
		</div>
		<login-button if={!authenticated} />
	</div>

	@user = null

	@on "mount", () ->
		# Retreive user data from the server
		qwest.post("/userdata").then ((resp) ->
			@user = JSON.parse resp
			@update()
		).bind this
</user-widget>

<login-button>

	<div class='widget' if={toggled}>
		<h4 class='title'>Login</h4>
		<form method='POST' action="/login" name="login-form">
			<input type='text'
			       name='name'
			       placeholder='Username'/>
			<br />
			<input type='password'
			       name='password'
			       placeholder='Password'/>
			<br />
			<button type='submit'>Sign in</button>
			<button type="button" onclick={toggle}>Cancel</button>
		</form>
	</div>
	<a href="#" if={!toggled} onclick={toggle}>Login</a>

	@toggled = false

	@toggle = (e) ->
		console.log e
		@toggled = !@toggled
		@update()

</login-button>