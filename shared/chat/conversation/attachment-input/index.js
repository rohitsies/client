// @noflow
import React, {Component} from 'react'
import {Box, Button, Icon, Input, PopupDialog, Text, ButtonBar} from '../../../common-adapters/index'
import {globalColors, globalStyles, isMobile} from '../../../styles'

type Props = {}
type State = {
  index: number,
  title: string,
}

class RenderAttachmentInput extends Component<Props, State> {
  state: State

  constructor(props: Props) {
    super(props)
    this.state = {
      index: 0,
      title: (props.inputs.length > 0 && props.inputs[0].title) || '',
    }
  }

  _onSelect = () => {
    const close = this.state.index === this.props.inputs.length - 1
    this.props.onSelect(this.props.inputs[this.state.index], this.state.title, close)
    if (!close) {
      const nextIndex = this.state.index + 1
      this.setState({
        index: nextIndex,
        title: this.props.inputs[nextIndex].title,
      })
    }
  }

  _updateTitle = title => {
    this.setState({title})
  }

  render() {
    const count = this.props.inputs.length
    const currentTitle =
      (this.props.inputs[this.state.index] && this.props.inputs[this.state.index].title) || ''

    return (
      <PopupDialog onClose={this.props.onClose}>
        <Box style={isMobile ? stylesMobile : stylesDesktop}>
          <Icon type="icon-file-uploading-48" />
          {count > 0 && (
            <Text type="BodySmall" style={{color: globalColors.black_40, marginTop: 5}}>
              {currentTitle} ({this.state.index + 1} of {count})
            </Text>
          )}
          <Input
            style={isMobile ? stylesInputMobile : stylesInputDesktop}
            autoFocus={true}
            floatingHintTextOverride="Title"
            value={this.state.title}
            onEnterKeyDown={this._onSelect}
            onChangeText={this._updateTitle}
          />
          <ButtonBar>
            <Button type="Secondary" onClick={this.props.onClose} label="Cancel" />
            <Button type="Primary" onClick={this._onSelect} label="Send" />
          </ButtonBar>
        </Box>
      </PopupDialog>
    )
  }
}

const stylesDesktop = {
  ...globalStyles.flexBoxColumn,
  alignItems: 'center',
  flex: 1,
  justifyContent: 'center',
  marginBottom: 80,
  marginLeft: 80,
  marginRight: 80,
  marginTop: 90,
}

const stylesInputDesktop = {
  marginTop: 70,
  width: 460,
}

const stylesMobile = {
  ...globalStyles.flexBoxColumn,
  alignItems: 'center',
  flex: 1,
  justifyContent: 'flex-start',
  marginTop: 40,
}

const stylesInputMobile = {
  marginTop: 40,
  minWidth: 320,
  paddingLeft: 20,
  paddingRight: 20,
}

export default RenderAttachmentInput
