kd = require 'kd'
KDNotificationView = kd.NotificationView
React = require 'react'
ReactView = require 'app/react/reactview'
getGroupStatus = require 'app/util/getGroupStatus'
whoami      = require 'app/util/whoami'
{ Status } = require 'app/redux/modules/payment/constants'

VerifyPasswordModal = require 'lab/VerifyPasswordModal'

Dialog = require 'lab/Dialog'
Modal = require 'lab/Modal'
Label = require 'lab/Text/Label'

module.exports = class VerifyPasswordModal2 extends ReactView


  renderReact: ->
    <VerifyPasswordModal
      onForgotPassword={@onForgotPassword.bind(this)}
      onPrimaryButtonClick={@onPrimaryButtonClick.bind(this)}
      onSecondaryButtonClick={@onSecondaryButtonClick.bind(this)} />

  onForgotPassword: ->
    account = whoami()
    account.fetchEmail (err, email) =>
      return @showError err  if err
      @doRecover email

  doRecover: (email) ->
    $.ajax
      url         : '/Recover'
      data        : { email, _csrf : Cookies.get '_csrf' }
      type        : 'POST'
      error       : (xhr) ->
        { responseText } = xhr
        new KDNotificationView { title : responseText }
      success     : ->
        new KDNotificationView
          title     : 'Check your email'
          content   : "We've sent you a password recovery code."
          duration  : 4500

  onPrimaryButtonClick: (e) ->
    console.log e

  onSecondaryButtonClick: ->
    console.log 'second'


    