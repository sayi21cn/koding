kd = require 'kd'
KDNotificationView = kd.NotificationView
React = require 'app/react'
ReactView = require 'app/react/reactview'
whoami = require 'app/util/whoami'
Link = require 'app/components/common/link'
SearchInputBox = require 'lab/SearchInputBox'

module.exports = class VerifiedPasswordModalContent extends ReactView

  forgotPassword: ->
    account = whoami()
    account.fetchEmail (err, email) =>
      return @showError err  if err
      @doRecover email

  onKeyUp: (event) ->
    @onIconClick()  if event.keyCode is 27

  onIconClick: (event) ->

    @setState
      close: yes
      searchQuery: ''

  handleChange: (event) ->
    @setState
      value: event.target.value
    console.log @state

  renderReact: ->

    hasTransfer = false

    <div>
      <div className="warning-prompt">You will not be able to access this team unless you are invited again. This action <span>cannot</span> be undone.</div>
      <main className="main-container">
        <p><strong>{if hasTransfer then '1- ' else ''}Please enter your password</strong> to continue:</p>
        <input className="kdinput text" type="password" placeholder="Enter password" onChange={@handleChange} />
        <Link onClick= { @bound 'forgotPassword' }>Forgot password?</Link>
        {@renderTransferOwnership() if hasTransfer}
      </main>
    </div>

