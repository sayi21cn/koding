React = require 'react'

module.exports = SearchInputBox = (props) ->

  { value, onChangeCallback, onFocusCallback, onKeyUp, active, placeholder } = props

  <input
    type='text'
    className="kdinput text searchStackInput#{if active then ' active' else ''}"
    placeholder={placeholder}
    value={value}
    onChange={onChangeCallback}
    onKeyUp={onKeyUp}
    onClick={onFocusCallback}
    onFocus={onFocusCallback} />

