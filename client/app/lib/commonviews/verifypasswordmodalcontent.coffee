kd = require 'kd'
React = require 'app/react'
ReactView = require 'app/react/reactview'
Link = require 'app/components/common/link'

module.exports = class VerifiedPasswordModalContent extends ReactView

  renderReact: ->
    <main className="main-container">
      <div className="warning-prompt">You will not be able to access this team unless you are invited again. This action <span>cannot</span> be undone.</div>
      <p><strong>1- Please enter your password</strong> to continue:</p>
      <input className="kdinput text" />
      <Link onClick= { @bound 'forgotPassword' }>Forgot password?</Link>
      <p><strong>2- Please transfer ownership to another member.</strong> Since you're the owner of this team, we need a new owner before you leave.</p>
      <input className="kdinput text" />
    </main>

  forgotPassword: ->
    console.log 'hello'