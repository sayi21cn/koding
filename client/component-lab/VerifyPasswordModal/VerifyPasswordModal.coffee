React = require 'react'
Modal = require 'lab/Modal'
Label = require 'lab/Text/Label'

styles = require './VerifyPasswordModal.stylus'

Input = require 'lab/Input/Input'
Link = require 'app/components/common/link'

module.exports = class VerifyPasswordModal extends React.Component

  constructor: ->
    super()
    @state = {a: ''}

  update: (e) ->
    @setState
      a: e.target.value

  render: ->
    { onForgotPassword, onPrimaryButtonClick, onSecondaryButtonClick } = @props

    hasTransfer = false

    modalProps =
      width: "xlarge"
      height: "normal"
      showAlien: no
      isOpen: yes
      contentLabel: "Verify Password"
      shouldCloseOnOverlayClick: yes

    <Modal {...modalProps}>
      <Modal.Header title="Leave Team" />
      <div>
        <div className={styles.warningPrompt}>You will not be able to access this team unless you are invited again. This action <span>cannot</span> be undone.</div>
        <main className="mainContainer">
          <p><strong>{if hasTransfer then '1- ' else ''}Please enter your password</strong> to continue:</p>

            <Input.Field
              mask={[/\d/, /\d/]}
              guide={off}
              name='exp_month'
              title='Expiration'
              placeholder='Enter Password' />
          <input className="kdinput text" type="password" placeholder="Enter password" onChange={@update.bind(this)} />
          <Link onClick={onForgotPassword}>Forgot password?</Link>
          {hasTransfer and <TransferOwnership />}
        </main>
      </div>
      <Modal.Footer
        primaryButtonTitle="Leave Team"
        secondaryButtonTitle="Cancel"
        onPrimaryButtonClick={@onPrimaryButtonClick.bind(this)}
        onSecondaryButtonClick={onSecondaryButtonClick} />
    </Modal>

  onPrimaryButtonClick: ->
    @props.onPrimaryButtonClick @state.a

TransferOwnership = ->

  <div>
    <hr />
    <p><strong>2- Please transfer ownership to another member.</strong> Since you're the owner of this team, we need a new owner before you leave.</p>
    <input className="kdinput text autocomplete" placeholder="Select a user" />
  </div>

    # <SearchInputBox
    #   placeholder='Select a user...'
    #   onKeyUp={ @bound 'onKeyUp' }
    #    />