debug = (require 'debug') 'nse:toolbar:banner'

kd = require 'kd'
JView = require 'app/jview'

Events = require '../events'


module.exports = class Banner extends JView


  constructor: (options = {}, data) ->

    options.cssClass = kd.utils.curry 'banner', options.cssClass

    data      ?=
      message  : 'Hello world!'
      action   :
        title  : 'Fix'
        event  : Events.Banner.ActionClicked
      _initial : yes

    super options, data

    @_wait = null

    @messageButton = new kd.ButtonView
      cssClass : 'message-button solid blue small'
      title    : @getData 'buttonTitle'

    @closeButton = new kd.ButtonView
      cssClass : 'close-button'
      callback : @bound 'close'


  setData: (data) ->

    super data

    { action, autohide, _initial } = @getData()

    unless data._initial

      kd.utils.killWait @_wait

      if action
        @messageButton.setTitle action.title  if action.title
        if typeof action.fn is 'function'
          @messageButton.setCallback action.fn
        else if actionEvent = action.event
          @messageButton.setCallback =>
            @emit Events.Banner.Close
            @emit Events.Action, actionEvent, (action.args ? [])...
        @messageButton.show()
      else
        @messageButton.hide()

      if autohide
        @_wait = kd.utils.wait autohide, @bound 'close'

    return data


  close: ->
    @emit Events.Banner.Close


  pistachio: ->
    '''
    {.message{#(message)}}{{> @messageButton}}{{> @closeButton}}
    '''
