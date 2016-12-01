React = require 'react'

module.exports = SearchInputBox = (props) ->

  { value, onChangeCallback, onFocusCallback, onKeyUp, active } = props

  <input
    type='text'
    className="kdinput text searchStackInput#{if active then ' active' else ''}"
    placeholder='Search Docs, AWS, S3, Azure, GCP...'
    value={value}
    onChange={onChangeCallback}
    onKeyUp={onKeyUp}
    onClick={onFocusCallback}
    onFocus={onFocusCallback} />

