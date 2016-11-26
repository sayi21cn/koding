kd = require 'kd'
React = require 'app/react'
ReactView = require 'app/react/reactview'

Modal = require 'lab/Modal'


class TestOutputModal extends ReactView
  constructor: (options = {}, data) ->
    options.title ?= ''
    options.isOpen ?= yes
    options.appendToDomBody ?= yes
    options.showTable ?= no
    super options
    @appendToDomBody()  if @getOptions().appendToDomBody

  renderFooter: ->
    <span>
      Still don’t know what’s wrong? <a href='//www.koding.com/docs' target='_blank'>Check the testing @ koding</a>
    </span>

  showTable: (event) ->
    console.log 'onClick', event
    @updateOptions { showTable: not @options.showTable }

  renderReact: ->

    buttonTitle = 'See Mocha Outputs'
    buttonTitle = 'See Failures'  unless @options.showTable

    <Modal width='xlarge' height='taller' showAlien={yes} isOpen={@options.isOpen}>
      <Modal.Header title={@options.title} />
      <Modal.Content>
        <div className='switch-button' onClick={@bound 'showTable'}>{buttonTitle}</div>
        <ShowTable showTable={@options.showTable} />
        <MochaOutput showTable={@options.showTable} />
      </Modal.Content>
      <Modal.TextFooter text={@renderFooter()} />
    </Modal>


MochaOutput = ({ showTable }) ->
  className = 'mocha-output'
  className = 'mocha-output hidden'  if showTable

  <div className={className}>
    <div id="mocha"></div>
  </div>

ShowTable = ({ showTable }) ->

  return <span />  unless showTable

  failureStore = kd.singletons.reactor.evaluate(['TestSuitesFailureStore'])
  fileNames = Object.keys failureStore

  return <span />  unless fileNames.length

  suiteResults = []
  [0..fileNames.length-1].forEach (i) ->
    fileName = fileNames[i]
    suiteResults.push <SuiteResult key={i} fileName={fileName} suites={failureStore[fileName]} />

  <div className='suite-results'>
   {suiteResults}
  </div>


SuiteResult = ({ suites, fileName }) ->

  return <span />  unless suites.length

  suiteItems = []
  suites.forEach (suite, i) ->
    { status, title } = suite
    suiteItems.push <SuiteResultItem key={i} title={title} status={status} />


  <div className='suite'>
    <label className='file-name'>FileName: </label><span>{fileName}</span>
    {suiteItems}
  </div>


SuiteResultItem = ({ title, status }) ->
  statusClassName = unless status is 'Hard To Implement' or status is 'Not Implemented' then 'error'
  else if status is 'Hard To Implement' then 'hti'
  else 'ni'

  <div className='suite-info-wrapper'>
    <label className='title'>Suite Name: </label><span>{title}</span>
    <label className='status'>Status: </label><span className={statusClassName}>{status}</span>
  </div>


module.exports = TestOutputModal