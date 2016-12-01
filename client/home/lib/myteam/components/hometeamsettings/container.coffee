kd              = require 'kd'
React           = require 'app/react'
TeamFlux        = require 'app/flux/teams'
KDReactorMixin  = require 'app/flux/base/reactormixin'
View            = require './view'
Encoder         = require 'htmlencode'
showError       = require 'app/util/showError'
ContentModal    = require 'app/components/contentModal'
globals         = require 'globals'

notify = (title, duration = 5000) -> new kd.NotificationView { title, duration }

module.exports = class HomeTeamSettingsContainer extends React.Component

  getDataBindings: ->
    return {
      team: TeamFlux.getters.team
    }


  constructor: (props) ->

    super props

    canEdit = kd.singletons.groupsController.canEditGroup()
    @state =
      canEdit: canEdit
      loading: no
      teamNameChanged: no


  componentDidMount: ->

    teamName = if @state.team? then Encoder.htmlDecode @state.team.get 'title' else ''
    @setState { teamName: teamName }


  onUploadInput: ->

    [file] = @refs.view.input.files

    return  unless file
    @setState { loading: yes }
    reader = new FileReader
    reader.onload = (reader, event) =>

      name = "#{@state.team.get 'slug'}-logo-#{Date.now()}.png"
      mimeType = file.type
      content = file

      TeamFlux.actions.uploads3({ name, content, mimeType })
      .then ({ url }) =>
        dataToUpdate =
          customize:
            logo: url
        @updateTeam { dataToUpdate }

      .timeout 15000

      .catch ({ err }) =>
        showError err  if err
        showError 'There was a problem while uploading your logo, please try again later!'  unless err
        @setState { loading: no }

    reader.readAsDataURL file


  onClickLogo: -> @refs.view.input.click()


  onRemoveLogo: ->
    @refs.view.input.value = null
    @updateTeam { dataToUpdate: customize: { logo: null } }


  onUpdate: ->

    return  if @state.teamName is @state.team.get 'title'

    title = Encoder.htmlEncode @state.teamName

    @updateTeam { dataToUpdate: { title } }


  updateTeam: ({ dataToUpdate }) ->

    TeamFlux.actions.updateTeam(dataToUpdate).then ({ message }) =>
      notify message
      @setState { loading: no, teamNameChanged: no }
    .catch ({ message }) =>
      notify message
      @setState { loading: no, teamNameChanged: no }


  onTeamNameChanged: (event) ->

    { value } = event.target

    @setState
      teamName: value
      teamNameChanged: @state.team.get('title') isnt value


  onLeaveTeam: (event) ->
    textIfOwner = if _globals.userRoles.indexOf('owner') != -1 then " We will transfer the ownership of these resources to team owner." else ""

    surePartial = "<div class='image'></div>
      <div class='caption'>
        <h2>You are about to leave the team Teamname.</h2>
        <p>Leaving this team will remove your access to your <strong>3 public stacks</strong>." + textIfOwner + " You can download these resources below to reuse them.</p>
        <a class='download-icon' href='#'>Download Stacks & Credentials</a>
      </div>
      <div class='clearfix'></div>"

    modal = new ContentModal
      width         : 639
      cssClass      : 'content-modal leave-team'
      title         : 'Are You Sure?'
      content       : surePartial
      overlay       : yes
      buttons       :
        'DontSave'  :
          cssClass  : 'solid cancel medium no-text-transform'
          title     : "Cancel"
          callback  : =>
            modal.destroy()
        'SaveClose' :
          cssClass  : 'solid medium no-text-transform'
          title     : 'Yes, Continue'
          callback  : =>
            modal.destroy()

            partial = '<div class="warning-prompt">You will not be able to access this team unless you are invited again. This action <span>cannot</span> be undone.</div>'

            TeamFlux.actions.leaveTeam(partial).catch (err) ->
              showError err

  render: ->

    <View
      ref='view'
      team={@state.team}
      teamName={@state.teamName}
      teamNameChanged={@state.teamNameChanged}
      canEdit={@state.canEdit}
      loading={@state.loading}
      onUploadInput={@bound 'onUploadInput'}
      onClickLogo={@bound 'onClickLogo'}
      onRemoveLogo={@bound 'onRemoveLogo'}
      onUpdate={@bound 'onUpdate'}
      onLeaveTeam={@bound 'onLeaveTeam'}
      onTeamNameChanged={@bound 'onTeamNameChanged'}/>




HomeTeamSettingsContainer.include [KDReactorMixin]
