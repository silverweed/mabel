window.TopBar = React.createClass {
        render: ->
                return (
                        <div id='topbar'>
                            <SearchBar source='/search'/>
                            <UserWidget source='/userdata'/>
                        </div>
                )
}
