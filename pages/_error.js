import React from 'react'
import Link from 'next/link'

class Error extends React.Component {
  static getInitialProps({ res, err }) {
    const statusCode = res ? res.statusCode : err ? err.statusCode : null
    return { statusCode }
  }

  render() {
    return (
      <p>
          Page not found, please return to <Link href="/">Home</Link>
      </p>
    )
  }
}

export default Error
