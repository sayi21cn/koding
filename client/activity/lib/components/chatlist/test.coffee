kd           = require 'kd'
React        = require 'kd-react'
ReactDOM     = require 'react-dom'
expect       = require 'expect'
TestUtils    = require 'react-addons-test-utils'
toImmutable  = require 'app/util/toImmutable'
ChatList     = require './index'
ChatListItem = require 'activity/components/chatlistitem'
DateMarker   = require 'activity/components/datemarker'

describe 'ChatList', ->

  messages = toImmutable [
    {
      id              : 1
      body            : 'Archaeologists develop new theories about how Stonehenge was built'
      interactions    : { like : { actorsCount : 1 } }
      repliesCount    : 2
      createdAt       : '2016-01-01'
      account         :
        _id           : 1
        profile       : { nickname : 'nick', firstName : '', lastName : '' }
        isIntegration : yes
    }
    {
      id              : 2
      body            : 'NASA is searching for the next generation of space explorers'
      interactions    : { like : { actorsCount : 3 } }
      repliesCount    : 5
      createdAt       : '2016-01-01'
      account         :
        _id           : 2
        profile       : { nickname : 'john', firstName : '', lastName : '' }
        isIntegration : yes
    }
    {
      id              : 3
      body            : 'A bike lane in the Netherlands has been designed to harness energy from the sun'
      interactions    : { like : { actorsCount : 2 } }
      repliesCount    : 3
      createdAt       : '2016-01-15'
      account         :
        _id           : 3
        profile       : { nickname : 'alex', firstName : '', lastName : '' }
        isIntegration : yes
    }
  ]

  describe '::render', ->

    it 'renders messages', ->

      result = TestUtils.renderIntoDocument(
        <ChatList messages={messages} showItemMenu=yes selectedMessageId={messages.last().get 'id'} />
      )
      items  = TestUtils.scryRenderedComponentsWithType result, ChatListItem

      expect(items.length).toEqual messages.size
      for item, i in items
        expect(item.props.message).toBe messages.get i
        expect(item.props.showItemMenu).toBe yes
      expect(items.last.props.isSelected).toBe yes


    it 'renders date markers', ->

      result = TestUtils.renderIntoDocument(
        <ChatList messages={messages} showItemMenu=yes selectedMessageId={messages.last().get 'id'} />
      )
      items  = TestUtils.scryRenderedComponentsWithType result, DateMarker

      expect(items.length).toEqual 2
      expect(items.first.props.date).toEqual messages.first().get 'createdAt'
      expect(items[1].props.date).toEqual messages.last().get 'createdAt'
