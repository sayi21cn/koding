kd = require 'kd'
KDModalViewWithForms = kd.ModalViewWithForms
KDNotificationView = kd.NotificationView
whoami = require 'app/util/whoami'
showError = require 'app/util/showError'
ContentModal = require 'app/components/contentModal'
VerifiedPasswordModalContent = require './verifypasswordmodalcontent'

module.exports = class VerifyPasswordModal extends ContentModal

  constructor: (buttonTitle = 'Submit', options = {}, callback) ->

    partial  = options.partial || ''
    cssClass = 'content-modal'
    cssClass = 'content-modal with-partial'  if partial

    options =
      title                       : options.title or 'Please verify your current password'
      cssClass                    : cssClass
      overlay                     : yes
      overlayClick                : no
      width                       : options.width or 605
      buttons               :
        Cancel              :
          style             : 'GenericButton cancel'
          title             : 'Cancel'
          callback          : => @destroy()
        Submit              :
          title             : buttonTitle
          style             : 'GenericButton'
          type              : 'submit'
      tabs                        :
        navigable                 : yes
        forms                     :
          verifyPasswordForm      :
            callback              : =>
              callback @modalTabs.forms.verifyPasswordForm.inputs.password.getValue()
              @destroy()
            fields                :
              planDetails     :
                type          : 'hidden'
                cssClass      : 'hidden'  unless partial
                nextElement   :
                  planDetails :
                    cssClass  : 'content'
                    itemClass : kd.View
                    partial   : partial
              password            :
                name              : 'password'
                cssClass          : 'line-with'
                label             : '<strong>Please enter your password</strong> to continue'
                placeholder       : 'Enter your current password'
                type              : 'password'
                validate          :
                  rules           :
                    required      : yes
                  messages        :
                    required      : 'Current Password required!'
                nextElement   :
                  planDetails :
                    cssClass  : 'content'
                    itemClass : kd.View
                    partial   : '<a href="#">Forgot password?</a>'
                    callback          : =>
                      account = whoami()
                      account.fetchEmail (err, email) =>
                        return @showError err  if err
                        @doRecover email
                        @destroy()
                  planDetailsss :
                    cssClass  : 'contentssdf'
                    itemClass : kd.View
                    partial   : '<a href="#">Forgot password?</a>'

    super options

  createBody: ->
    @addSubView new VerifiedPasswordModalContent()




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
