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
                                <button type='submit'>Log out</button>
                            </form>
                        </div>
                else
                        <div>
                            <LoginButton />
                        </div>
}

LoginButton = React.createClass {
        getInitialState: ->
                { toggled: no }

        render: ->
                if @state.toggled
                        <div className='widget'>
                            <h4 className='title'>Login</h4>
                            <form method='POST' action="/login" name="login-form">
                                <input type='text'
                                       name='name'
                                       placeholder='Username'/>
                                <br/>
                                <input type='password'
                                       name='password'
                                       placeholder='Password'/>
                                <br/>
                                <button type='submit'>Sign in</button>
                                <button type="button" onClick={@toggle}>Cancel</button>
                            </form>
                        </div>
                else
                        <a href="#" onClick={@toggle}>Login</a>
        toggle: ->
                @setState { toggled: !@state.toggled }
}