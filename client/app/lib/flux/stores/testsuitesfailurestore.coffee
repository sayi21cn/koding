actions = require '../actions/actiontypes'
KodingFluxStore = require 'app/flux/base/store'

module.exports = class TestSuitesFailureStore extends KodingFluxStore

  @getterPath = 'TestSuitesFailureStore'

  getInitialState: -> {}

  initialize: ->
    @on actions.TEST_SUITE_FAIL, @add

  add: (state, { title, status, parentTitle }) ->
    unless status is 'Hard To Implement' or status is 'Not Implemented'
      status = 'Error'
    state[parentTitle] ?= []
    state[parentTitle].push { title, status }
    return state


