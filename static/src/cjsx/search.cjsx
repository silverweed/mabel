window.SearchBar = React.createClass {
        render: ->
                return (
                        <form className='searchBar' method='POST' action={@props.source}>
                            <input type='text' placeholder='Search by tag'/>
                            <button type='submit'>Search</button>
                        </form>
                )
}
