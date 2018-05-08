// @flow
import React from 'react'
import {Text} from '../common-adapters'
import {storiesOf} from '../stories/storybook'
import asset from './asset/index.stories'
import transaction from './transaction/index.stories'
import walletList from './wallet-list/index.stories'
import wallet from './wallet/index.stories'
import sendForm from './send-form/index.stories'

const load = () => {
  asset()
  sendForm()
  walletList()
  wallet()
  transaction()

  /* Still TODO */
  storiesOf('Wallets', module).add('Wallet Onboarding', () => (
    <Text type="BodyBig">Wallet Onboarding TBD</Text>
  ))
  storiesOf('Wallets', module).add('Settings', () => <Text type="BodyBig">Settings TBD</Text>)
}

export default load
