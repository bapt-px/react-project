import Link from 'next/link'
import React, { useState } from 'react';

export default () => {

  const [url, setUrl] = useState('')
  const [nbComment, setNbComment] = useState(0)
  const [average, setAverage] = useState(0)

  const searchPage = e => {
    fetch('http://www.localhost:80/test', { method: 'GET',
      headers: new Headers(),
      mode: 'no-cors',
      cache: 'default' })
      .then(response => {
        setNbComment(5)
        setAverage(5)
      })
  }

  return (
    <div>
      <input type="text" name="url" value={url} onChange={e => setUrl(e.target.value)} /><br/>
      <button onClick={() => searchPage()}>
        Click me
      </button>
      <p>
        Words: {nbComment}<br/>
        Average:: {average}
      </p>

  </div>
  )
}
