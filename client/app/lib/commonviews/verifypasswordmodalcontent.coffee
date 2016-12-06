kd = require 'kd'
React = require 'app/react'
ReactView = require 'app/react/reactview'
Link = require 'app/components/common/link'
SearchInputBox = require 'lab/SearchInputBox'

module.exports = class VerifiedPasswordModalContent extends ReactView

  forgotPassword: ->
    console.log 'hello'

  onKeyUp: (event) ->
    @onIconClick()  if event.keyCode is 27

  onIconClick: (event) ->

    @setState
      close: yes
      searchQuery: ''

  renderReact: ->
    <div>
      <div className="warning-prompt">You will not be able to access this team unless you are invited again. This action <span>cannot</span> be undone.</div>
      <main className="main-container">
        <p><strong>1- Please enter your password</strong> to continue:</p>
        <input className="kdinput text" type="password" placeholder="Enter password" />
        <Link onClick= { @bound 'forgotPassword' }>Forgot password?</Link>
        <hr />
        <p><strong>2- Please transfer ownership to another member.</strong> Since you're the owner of this team, we need a new owner before you leave.</p>
        <input className="kdinput text autocomplete" placeholder="Select a user" />
        <SearchInputBox
          placeholder='Select a user...'
          onKeyUp={ @bound 'onKeyUp' }
           />
      </main>
    </div>
