// @flow
import * as React from 'react'
import {Avatar, Box2, Text} from '../../common-adapters'
import {globalMargins, styleSheetCreate} from '../../styles'
// TODO: Format relative dates.
import {formatTimeForPopup} from '../../util/timestamp'

export type Props = {
  sourceValue: string,
  targetValue: string,
  sender: string,
  receiver: string,
  note: string,
  timestamp: Date,
}

export const Transaction = (props: Props) => {
  return (
    <Box2 direction="horizontal" fullWidth={true} style={styles.container}>
      <Avatar username={props.sender} size={48} />
      <Box2 direction="vertical" fullHeight={true} fullWidth={true} style={styles.rightContainer}>
        <Text type="BodySmall">{formatTimeForPopup(props.timestamp)}</Text>
        <Box2 direction="horizontal" fullHeight={true} fullWidth={true} style={styles.rightDownContainer}>
          <Box2 direction="vertical" fullHeight={true} fullWidth={true} style={styles.detailContainer}>
            <Text type="BodySemibold" lineClamp={1}>
              {props.sender} sent Lumens worth {props.sourceValue} to {props.receiver}
            </Text>
            <Text type="Body">{props.note}</Text>
          </Box2>
          <Text type="BodySmall" lineClamp={1}>
            {props.targetValue}
          </Text>
        </Box2>
      </Box2>
    </Box2>
  )
}

const styles = styleSheetCreate({
  container: {
    border: 'solid 1px black',
    minHeight: 80,
    padding: globalMargins.tiny,
    paddingRight: globalMargins.small,
  },
  detailContainer: {
    border: 'solid 1px black',
  },
  rightContainer: {
    border: 'solid 1px black',
    marginLeft: globalMargins.tiny,
  },
  rightDownContainer: {
    border: 'solid 1px black',
  },
})

export default Transaction
